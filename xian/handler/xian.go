package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	xian "xian/proto/xian"
)

type Xian struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Xian) Call(ctx context.Context, req *xian.Request, rsp *xian.Response) error {
	log.Info("Received Xian.Call request")
	rsp.Msg = "Hello, China No1, " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Xian) Stream(ctx context.Context, req *xian.StreamingRequest, stream xian.Xian_StreamStream) error {
	log.Infof("Received Xian.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&xian.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Xian) PingPong(ctx context.Context, stream xian.Xian_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&xian.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
