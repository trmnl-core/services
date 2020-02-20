package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	elephant "elephant/proto/elephant"
)

type Elephant struct{}

func (e *Elephant) Handle(ctx context.Context, msg *elephant.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *elephant.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
