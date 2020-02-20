package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	whale "whale/proto/whale"
)

type Whale struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Whale) Call(ctx context.Context, req *whale.Request, rsp *whale.Response) error {
	log.Log("Received Whale.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Whale) Stream(ctx context.Context, req *whale.StreamingRequest, stream whale.Whale_StreamStream) error {
	log.Logf("Received Whale.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&whale.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Whale) PingPong(ctx context.Context, stream whale.Whale_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&whale.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
