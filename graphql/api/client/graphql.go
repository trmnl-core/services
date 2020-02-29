package client

import (
	"context"

	graphql "api/proto/graphql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type graphqlKey struct{}

// FromContext retrieves the client from the Context
func GraphqlFromContext(ctx context.Context) (graphql.GraphqlService, bool) {
	c, ok := ctx.Value(graphqlKey{}).(graphql.GraphqlService)
	return c, ok
}

// Client returns a wrapper for the GraphqlClient
func GraphqlWrapper(service micro.Service) server.HandlerWrapper {
	client := graphql.NewGraphqlService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, graphqlKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
