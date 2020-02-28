package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	rex "rex-srv/proto/rex"
)

type Rex struct{}

func (e *Rex) Handle(ctx context.Context, msg *rex.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *rex.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
