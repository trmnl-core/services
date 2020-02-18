package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	test2 "test2/proto/test2"
)

type Test2 struct{}

func (e *Test2) Handle(ctx context.Context, msg *test2.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test2.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
