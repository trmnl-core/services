package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	johnny "johnny-test1/proto/johnny"
)

type Johnny struct{}

func (e *Johnny) Handle(ctx context.Context, msg *johnny.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *johnny.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
