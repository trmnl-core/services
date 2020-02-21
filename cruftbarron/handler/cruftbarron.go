package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	cruftbarron "cruftbarron/proto/cruftbarron"
)

type Cruftbarron struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Cruftbarron) Call(ctx context.Context, req *cruftbarron.Request, rsp *cruftbarron.Response) error {
	log.Log("Received Cruftbarron.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Cruftbarron) Stream(ctx context.Context, req *cruftbarron.StreamingRequest, stream cruftbarron.Cruftbarron_StreamStream) error {
	log.Logf("Received Cruftbarron.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&cruftbarron.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Cruftbarron) PingPong(ctx context.Context, stream cruftbarron.Cruftbarron_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&cruftbarron.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
