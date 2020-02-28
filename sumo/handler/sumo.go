package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	sumo "sumo/proto/sumo"
)

type Sumo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Sumo) Call(ctx context.Context, req *sumo.Request, rsp *sumo.Response) error {
	log.Info("Received v2 Sumo.Call request")
	rsp.Msg = "Hello U:" + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Sumo) Stream(ctx context.Context, req *sumo.StreamingRequest, stream sumo.Sumo_StreamStream) error {
	log.Infof("Received v2 Sumo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&sumo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Sumo) PingPong(ctx context.Context, stream sumo.Sumo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&sumo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
