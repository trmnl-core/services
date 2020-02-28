package handler

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/services/location/ingester"
	loc "github.com/micro/services/location/proto/location"

	"golang.org/x/net/context"
)

func (l *Location) Save(ctx context.Context, req *loc.SaveRequest, rsp *loc.SaveResponse) error {
	log.Print("Received Location.Save request")

	entity := req.GetEntity()

	if entity.GetLocation() == nil {
		return errors.BadRequest(server.DefaultOptions().Name+".save", "Require location")
	}

	p := micro.NewPublisher(ingester.Topic, client.DefaultClient)
	if err := p.Publish(ctx, entity); err != nil {
		return errors.InternalServerError(server.DefaultOptions().Name+".save", err.Error())
	}

	return nil
}
