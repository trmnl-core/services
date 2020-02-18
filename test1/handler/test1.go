package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	test1 "test1/proto/test1"
)

type Test1 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test1) Call(ctx context.Context, req *test1.Request, rsp *test1.Response) error {
	log.Log("Received Test1.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test1) Stream(ctx context.Context, req *test1.StreamingRequest, stream test1.Test1_StreamStream) error {
	log.Logf("Received Test1.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test1.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test1) PingPong(ctx context.Context, stream test1.Test1_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test1.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
