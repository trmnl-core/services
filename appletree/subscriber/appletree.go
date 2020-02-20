package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	appletree "appletree/proto/appletree"
)

type Appletree struct{}

func (e *Appletree) Handle(ctx context.Context, msg *appletree.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *appletree.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
