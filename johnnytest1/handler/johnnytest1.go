package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	johnnytest1 "johnnytest1/proto/johnnytest1"
)

type Johnnytest1 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Johnnytest1) Call(ctx context.Context, req *johnnytest1.Request, rsp *johnnytest1.Response) error {
	log.Log("Received Johnnytest1.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Johnnytest1) Stream(ctx context.Context, req *johnnytest1.StreamingRequest, stream johnnytest1.Johnnytest1_StreamStream) error {
	log.Logf("Received Johnnytest1.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&johnnytest1.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Johnnytest1) PingPong(ctx context.Context, stream johnnytest1.Johnnytest1_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&johnnytest1.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
