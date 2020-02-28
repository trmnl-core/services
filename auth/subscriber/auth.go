package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	auth "auth/proto/auth"
)

type Auth struct{}

func (e *Auth) Handle(ctx context.Context, msg *auth.User) error {
	log.Info("Handler Received message: ", msg.Firstname)
	return nil
}

func Handler(ctx context.Context, msg *auth.User) error {
	log.Info("Function Received message: ", msg.Firstname)
	return nil
}
