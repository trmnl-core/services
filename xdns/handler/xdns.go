package handler

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/OWASP/Amass/v3/config"
	"github.com/OWASP/Amass/v3/datasrcs"
	"github.com/OWASP/Amass/v3/enum"
	"github.com/OWASP/Amass/v3/systems"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/micro/v3/service"
	mconfig "github.com/micro/micro/v3/service/config"
	"github.com/patrickmn/go-cache"

	xdns "github.com/trmnl-core/services/xdns/proto/xdns"

	aproto "github.com/trmnl-core/services/alert/proto/alert"

	mstore "github.com/micro/micro/v3/service/store"
)

type XDNS struct {
	alertService aproto.AlertService
	auth         auth.Auth
	config       conf
	cache        *cache.Cache
}

type conf struct {
	TestMode bool `json:"test_env"`
}

func NewXDNS(srv *service.Service, auth auth.Auth) *XDNS {
	rand.Seed(time.Now().UTC().UnixNano())

	c := conf{}
	val, err := mconfig.Get("trmnl.xdns")
	if err != nil {
		logger.Warnf("Error getting config: %v", err)
	}

	err = val.Scan(&c)
	if err != nil {
		logger.Warnf("Error scanning config: %v", err)
	}

	s := &XDNS{
		auth:         auth,
		config:       c,
		cache:        cache.New(1*time.Minute, 5*time.Minute),
		alertService: aproto.NewAlertService("alert", srv.Client()),
	}
	return s
}

func (x *XDNS) EnumerateDNS(ctx context.Context,
	req *xdns.EnumerateDNSRequest,
	rsp *xdns.EnumerateDNSResponse) error {

	return nil
}

func enumerateDNS(req *xdns.EnumerateDNSRequest) error {
	if req.Domain == "" {
		return errors.New("missing domain")
	}
	settings := req.Settings
	if settings == nil {
		return errors.New("missing settings")
	}

	cfg := config.NewConfig()
	cfg.AddDomain(req.Domain)

	sys, err := systems.NewLocalSystem(cfg)
	if err != nil {
		return err
	}
	sys.SetDataSources(datasrcs.GetAllSources(sys))
	e := enum.NewEnumeration(cfg, sys)
	if e == nil {
		return
	}
	defer e.Close()

	e.Start()
	results := e.ExtractOutput(nil)
	b, err := json.Marshal(results)
	if err != nil {
		return err
	}
	if err := mstore.Write(&mstore.Record{
		Key:   req.Domain,
		Value: b,
	}); err != nil {
		return err
	}
	return nil
}

func (x *XDNS) Status(ctx context.Context,
	req *xdns.StatusRequest,
	rsp *xdns.StatusResponse) error {
	return nil
}
