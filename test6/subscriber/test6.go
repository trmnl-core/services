package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	test6 "test6/proto/test6"
)

type Test6 struct{}

func (e *Test6) Handle(ctx context.Context, msg *test6.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test6.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
