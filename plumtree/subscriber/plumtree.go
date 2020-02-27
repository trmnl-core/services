package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	plumtree "plumtree/proto/plumtree"
)

type Plumtree struct{}

func (e *Plumtree) Handle(ctx context.Context, msg *plumtree.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *plumtree.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
