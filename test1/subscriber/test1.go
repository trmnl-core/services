package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	test1 "test1/proto/test1"
)

type Test1 struct{}

func (e *Test1) Handle(ctx context.Context, msg *test1.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test1.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
