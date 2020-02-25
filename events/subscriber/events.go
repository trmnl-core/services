package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	events "events/proto/events"
)

type Events struct{}

func (e *Events) Handle(ctx context.Context, msg *events.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *events.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
