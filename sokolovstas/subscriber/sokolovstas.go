package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	sokolovstas "sokolovstas/proto/sokolovstas"
)

type Sokolovstas struct{}

func (e *Sokolovstas) Handle(ctx context.Context, msg *sokolovstas.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *sokolovstas.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
