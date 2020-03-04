package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	sokolovstas "sokolovstas/proto/sokolovstas"
)

type Sokolovstas struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Sokolovstas) Call(ctx context.Context, req *sokolovstas.Request, rsp *sokolovstas.Response) error {
	log.Info("Received Sokolovstas.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Sokolovstas) Stream(ctx context.Context, req *sokolovstas.StreamingRequest, stream sokolovstas.Sokolovstas_StreamStream) error {
	log.Infof("Received Sokolovstas.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&sokolovstas.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Sokolovstas) PingPong(ctx context.Context, stream sokolovstas.Sokolovstas_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&sokolovstas.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
