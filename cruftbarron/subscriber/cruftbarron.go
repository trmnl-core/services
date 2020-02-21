package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	cruftbarron "cruftbarron/proto/cruftbarron"
)

type Cruftbarron struct{}

func (e *Cruftbarron) Handle(ctx context.Context, msg *cruftbarron.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *cruftbarron.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
