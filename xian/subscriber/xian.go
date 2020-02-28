package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	xian "xian/proto/xian"
)

type Xian struct{}

func (e *Xian) Handle(ctx context.Context, msg *xian.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *xian.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
