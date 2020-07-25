package subscriber

import (
	"context"

	nproto "github.com/micro/services/notifications/service/proto/notifications"
)

type Subscriber struct {
}

func (s *Subscriber) Handle(ctx context.Context, msg *nproto.Event) error {
	return nil
}
