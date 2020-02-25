package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	events "events/proto/events"
)

type Events struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Events) Call(ctx context.Context, req *events.Request, rsp *events.Response) error {
	log.Log("Received Events.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Events) Stream(ctx context.Context, req *events.StreamingRequest, stream events.Events_StreamStream) error {
	log.Logf("Received Events.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&events.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Events) PingPong(ctx context.Context, stream events.Events_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&events.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
