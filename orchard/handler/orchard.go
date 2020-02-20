package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	orchard "orchard/proto/orchard"
)

type Orchard struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Orchard) Call(ctx context.Context, req *orchard.Request, rsp *orchard.Response) error {
	log.Log("Received Orchard.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Orchard) Stream(ctx context.Context, req *orchard.StreamingRequest, stream orchard.Orchard_StreamStream) error {
	log.Logf("Received Orchard.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&orchard.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Orchard) PingPong(ctx context.Context, stream orchard.Orchard_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&orchard.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
