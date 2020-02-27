package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	plumtree "plumtree/proto/plumtree"
)

type Plumtree struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Plumtree) Call(ctx context.Context, req *plumtree.Request, rsp *plumtree.Response) error {
	log.Info("Received Plumtree.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Plumtree) Stream(ctx context.Context, req *plumtree.StreamingRequest, stream plumtree.Plumtree_StreamStream) error {
	log.Infof("Received Plumtree.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&plumtree.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Plumtree) PingPong(ctx context.Context, stream plumtree.Plumtree_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&plumtree.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
