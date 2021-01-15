package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service/config"
	"github.com/slack-go/slack"

	usage "github.com/trmnl-core/services/usage/proto"

	pb "github.com/micro/micro/v3/proto/auth"
	rproto "github.com/micro/micro/v3/proto/runtime"
	"github.com/micro/micro/v3/service/client"
	log "github.com/micro/micro/v3/service/logger"
	nsproto "github.com/trmnl-core/services/namespaces/proto"
)

const (
	defaultNamespace = "micro"
)

type Usage struct {
	ns       nsproto.NamespacesService
	as       pb.AccountsService
	runtime  rproto.RuntimeService
	slackbot *slack.Client
}

func NewUsage(ns nsproto.NamespacesService, as pb.AccountsService, runtime rproto.RuntimeService) *Usage {
	val, err := config.Get("trmnl.alert.slack_token")
	if err != nil {
		log.Warnf("Error getting config: %v", err)
	}
	slackToken := val.String("")
	if len(slackToken) == 0 {
		log.Fatal("Missing required config micro.alert.slack_token")
	}

	u := &Usage{
		ns:       ns,
		as:       as,
		runtime:  runtime,
		slackbot: slack.New(slackToken),
	}
	return u
}

// Read account history by namespace, or lists latest values for each namespace if history is not provided.
func (e *Usage) Read(ctx context.Context, req *usage.ReadRequest, rsp *usage.ReadResponse) error {
	log.Infof("Received Usage.Read request, reading namespace '%v'", req.Namespace)
	u, err := e.usageForNamespace(req.Namespace)
	if err != nil {
		return err
	}
	rsp.Accounts = []*usage.Account{
		{
			Namespace: req.Namespace,
			Users:     u.Users,
			Services:  u.Services,
		},
	}
	return nil
}

type usg struct {
	Users     int64
	Services  int64
	Namespace string
}

func (e *Usage) usageForNamespace(namespace string) (*usg, error) {
	arsp, err := e.as.List(context.TODO(), &pb.ListAccountsRequest{
		Options: &pb.Options{
			Namespace: namespace,
		},
	}, client.WithAuthToken())
	if err != nil {
		return nil, err
	}
	userCount := 0
	for _, account := range arsp.Accounts {
		if account.Type == "user" {
			userCount++
		}
	}
	rrsp, err := e.runtime.Read(context.TODO(), &rproto.ReadRequest{
		Options: &rproto.ReadOptions{
			Namespace: namespace,
		},
	}, client.WithAuthToken(), client.WithRequestTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}
	serviceCount := len(rrsp.Services)
	return &usg{
		Users:     int64(userCount),
		Services:  int64(serviceCount),
		Namespace: namespace,
	}, nil
}

func (e *Usage) List(ctx context.Context, request *usage.ListRequest, response *usage.ListResponse) error {
	usages, err := e.usageForAllNamespaces()
	if err != nil {
		return err
	}

	response.Accounts = make([]*usage.Account, len(usages))
	for i, us := range usages {
		response.Accounts[i] = &usage.Account{
			Namespace: us.Namespace,
			Users:     us.Users,
			Services:  us.Services,
		}
	}

	nsCount := len(usages)
	var svcCount, userCount int64
	for _, u := range usages {
		svcCount += u.Services
		userCount += u.Users
	}

	response.Summary = &usage.Summary{
		NamespaceCount: int64(nsCount),
		UserCount:      userCount,
		ServicesCount:  svcCount,
	}
	return nil
}

func (e *Usage) CheckUsageCron() {
	usages, err := e.usageForAllNamespaces()
	if err != nil {
		log.Errorf("Error calculating usage %s", err)
		return
	}
	nsCount := len(usages)
	var svcCount, userCount int64
	for _, u := range usages {
		svcCount += u.Services
		userCount += u.Users
	}
	msg := fmt.Sprintf("Usage summary\nNamespaces: %d\nUsers: %d\nServices: %d\nFor a more detailed breakdown run `micro usage list`", nsCount, userCount, svcCount)

	valUser, _ := config.Get("trmnl.usage.cron.user")
	valChan, _ := config.Get("trmnl.usage.cron.channel")
	e.slackbot.SendMessage(valChan.String("team-important"),
		slack.MsgOptionUsername(valUser.String("Usage Service")),
		slack.MsgOptionText(msg, false),
	)

}

func (e *Usage) usageForAllNamespaces() ([]*usg, error) {
	rsp, err := e.ns.List(context.TODO(), &nsproto.ListRequest{}, client.WithAuthToken())
	if err != nil {
		return nil, err
	}
	jobs := make(chan string, 5)
	res := make(chan *usg, 5)
	defer close(res)
	defer close(jobs)
	errMap := map[string]int{}
	worker := func() {
		for nsID := range jobs {
			usg, err := e.usageForNamespace(nsID)
			if err != nil {
				if errMap[nsID] < 3 {
					errMap[nsID] += 1
					// put it back on
					jobs <- nsID
					continue
				}
				// too many errors, break out
				usg.Namespace = nsID
				log.Errorf("Too many errors for %s", nsID)
			}
			res <- usg
		}

	}

	//worker pool
	for i := 0; i < 3; i++ {
		go worker()
	}

	usages := []*usg{}
	go func() {
		for _, ns := range rsp.Namespaces {
			jobs <- ns.Id
		}

	}()
	for i := 0; i < len(rsp.Namespaces); i++ {
		usages = append(usages, <-res)
	}
	return usages, nil

}
