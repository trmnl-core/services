package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	barrel "barrel/proto/barrel"
)

type Barrel struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Barrel) Call(ctx context.Context, req *barrel.Request, rsp *barrel.Response) error {
	log.Log("Received Barrel.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Barrel) Stream(ctx context.Context, req *barrel.StreamingRequest, stream barrel.Barrel_StreamStream) error {
	log.Logf("Received Barrel.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&barrel.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Barrel) PingPong(ctx context.Context, stream barrel.Barrel_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&barrel.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
