package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	auth "auth/proto/auth"
)

type Auth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.LoginResponse) error {
	log.Info("Received Auth.Login request")
	rsp.Token = "my-sweet-token"
	rsp.User = &auth.User{
		Id:        "0000001-asdf",
		Firstname: "Rosy",
		Lastname:  "Rex",
		Email:     req.Email,
	}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Auth) Register(ctx context.Context, req *auth.RegisterRequest,  rsp *auth.RegisterResponse) error {
	log.Infof("Received Auth.Register")
	rsp.Message = "Hurray!!! " + req.Firstname + " You have registered successfully"

	return nil
}