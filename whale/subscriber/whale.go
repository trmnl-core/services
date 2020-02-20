package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	whale "whale/proto/whale"
)

type Whale struct{}

func (e *Whale) Handle(ctx context.Context, msg *whale.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *whale.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
