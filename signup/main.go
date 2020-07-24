package main

import (
	"github.com/m3o/services/signup/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	inviteproto "github.com/m3o/services/account/invite/proto"
	k8sproto "github.com/m3o/services/kubernetes/service/proto"
	paymentsproto "github.com/m3o/services/payments/provider/proto"
	signup "github.com/m3o/services/signup/proto/signup"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.signup"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	signup.RegisterSignupHandler(service.Server(), handler.NewSignup(
		paymentsproto.NewProviderService("go.micro.service.payment.stripe", service.Options().Client),
		inviteproto.NewInviteService("go.micro.service.account.invite", service.Options().Client),
		k8sproto.NewKubernetesService("go.micro.service.kubernetes", service.Options().Client),
		service.Options().Store,
		service.Options().Config,
		service.Options().Auth,
	))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
