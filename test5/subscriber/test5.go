package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	test5 "test5/proto/test5"
)

type Test5 struct{}

func (e *Test5) Handle(ctx context.Context, msg *test5.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test5.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
