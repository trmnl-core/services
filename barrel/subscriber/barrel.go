package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	barrel "barrel/proto/barrel"
)

type Barrel struct{}

func (e *Barrel) Handle(ctx context.Context, msg *barrel.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *barrel.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
