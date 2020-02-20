package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	elephant "elephant/proto/elephant"
)

type Elephant struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Elephant) Call(ctx context.Context, req *elephant.Request, rsp *elephant.Response) error {
	log.Log("Received Elephant.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Elephant) Stream(ctx context.Context, req *elephant.StreamingRequest, stream elephant.Elephant_StreamStream) error {
	log.Logf("Received Elephant.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&elephant.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Elephant) PingPong(ctx context.Context, stream elephant.Elephant_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&elephant.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
