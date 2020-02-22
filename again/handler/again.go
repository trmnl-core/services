package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	again "again/proto/again"
)

type Again struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Again) Call(ctx context.Context, req *again.Request, rsp *again.Response) error {
	log.Log("Received Again.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Again) Stream(ctx context.Context, req *again.StreamingRequest, stream again.Again_StreamStream) error {
	log.Logf("Received Again.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&again.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Again) PingPong(ctx context.Context, stream again.Again_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&again.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
