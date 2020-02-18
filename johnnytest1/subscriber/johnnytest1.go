package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	johnnytest1 "johnnytest1/proto/johnnytest1"
)

type Johnnytest1 struct{}

func (e *Johnnytest1) Handle(ctx context.Context, msg *johnnytest1.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *johnnytest1.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
