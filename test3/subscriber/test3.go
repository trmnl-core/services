package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	test3 "test3/proto/test3"
)

type Test3 struct{}

func (e *Test3) Handle(ctx context.Context, msg *test3.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test3.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
