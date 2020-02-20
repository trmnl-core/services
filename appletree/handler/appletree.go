package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	appletree "appletree/proto/appletree"
)

type Appletree struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Appletree) Call(ctx context.Context, req *appletree.Request, rsp *appletree.Response) error {
	log.Log("Received Appletree.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Appletree) Stream(ctx context.Context, req *appletree.StreamingRequest, stream appletree.Appletree_StreamStream) error {
	log.Logf("Received Appletree.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&appletree.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Appletree) PingPong(ctx context.Context, stream appletree.Appletree_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&appletree.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
