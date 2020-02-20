package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	orchard "orchard/proto/orchard"
)

type Orchard struct{}

func (e *Orchard) Handle(ctx context.Context, msg *orchard.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *orchard.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
