package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	sumo "sumo/proto/sumo"
)

type Sumo struct{}

func (e *Sumo) Handle(ctx context.Context, msg *sumo.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *sumo.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
