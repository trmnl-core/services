package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	test3 "test3/proto/test3"
)

type Test3 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test3) Call(ctx context.Context, req *test3.Request, rsp *test3.Response) error {
	log.Log("Received Test3.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test3) Stream(ctx context.Context, req *test3.StreamingRequest, stream test3.Test3_StreamStream) error {
	log.Logf("Received Test3.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test3.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test3) PingPong(ctx context.Context, stream test3.Test3_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test3.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
