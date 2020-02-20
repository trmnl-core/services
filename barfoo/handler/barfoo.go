package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	barfoo "barfoo/proto/barfoo"
)

type Barfoo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Barfoo) Call(ctx context.Context, req *barfoo.Request, rsp *barfoo.Response) error {
	log.Log("Received Barfoo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Barfoo) Stream(ctx context.Context, req *barfoo.StreamingRequest, stream barfoo.Barfoo_StreamStream) error {
	log.Logf("Received Barfoo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&barfoo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Barfoo) PingPong(ctx context.Context, stream barfoo.Barfoo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&barfoo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
