package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	test5 "test5/proto/test5"
)

type Test5 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test5) Call(ctx context.Context, req *test5.Request, rsp *test5.Response) error {
	log.Log("Received Test5.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test5) Stream(ctx context.Context, req *test5.StreamingRequest, stream test5.Test5_StreamStream) error {
	log.Logf("Received Test5.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test5.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test5) PingPong(ctx context.Context, stream test5.Test5_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test5.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
