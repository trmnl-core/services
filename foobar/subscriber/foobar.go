package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	foobar "foobar/proto/foobar"
)

type Foobar struct{}

func (e *Foobar) Handle(ctx context.Context, msg *foobar.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *foobar.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
