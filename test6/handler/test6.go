package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	test6 "test6/proto/test6"
)

type Test6 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test6) Call(ctx context.Context, req *test6.Request, rsp *test6.Response) error {
	log.Log("Received Test6.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test6) Stream(ctx context.Context, req *test6.StreamingRequest, stream test6.Test6_StreamStream) error {
	log.Logf("Received Test6.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test6.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test6) PingPong(ctx context.Context, stream test6.Test6_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test6.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
