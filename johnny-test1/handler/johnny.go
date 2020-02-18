package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	johnny "johnny-test1/proto/johnny"
)

type Johnny struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Johnny) Call(ctx context.Context, req *johnny.Request, rsp *johnny.Response) error {
	log.Log("Received Johnny.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Johnny) Stream(ctx context.Context, req *johnny.StreamingRequest, stream johnny.Johnny_StreamStream) error {
	log.Logf("Received Johnny.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&johnny.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Johnny) PingPong(ctx context.Context, stream johnny.Johnny_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&johnny.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
