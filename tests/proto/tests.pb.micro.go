// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/tests.proto

package go_micro_service_tests

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Tests service

func NewTestsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Tests service

type TestsService interface {
	// enables registering an endpoint for callback to run tests
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
}

type testsService struct {
	c    client.Client
	name string
}

func NewTestsService(name string, c client.Client) TestsService {
	return &testsService{
		c:    c,
		name: name,
	}
}

func (c *testsService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Tests.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tests service

type TestsHandler interface {
	// enables registering an endpoint for callback to run tests
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
}

func RegisterTestsHandler(s server.Server, hdlr TestsHandler, opts ...server.HandlerOption) error {
	type tests interface {
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
	}
	type Tests struct {
		tests
	}
	h := &testsHandler{hdlr}
	return s.Handle(s.NewHandler(&Tests{h}, opts...))
}

type testsHandler struct {
	TestsHandler
}

func (h *testsHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.TestsHandler.Register(ctx, in, out)
}
