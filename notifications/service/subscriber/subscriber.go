package subscriber

import (
	"context"

	nproto "github.com/m3o/services/notifications/service/proto/notifications"
)

type Subscriber struct {
}

func (s *Subscriber) Handle(ctx context.Context, msg *nproto.Event) error {
	return nil
}
