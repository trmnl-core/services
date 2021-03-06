// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/platform.proto

package platform

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

// Api Endpoints for Platform service

func NewPlatformEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Platform service

type PlatformService interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...client.CallOption) (*CreateNamespaceResponse, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...client.CallOption) (*DeleteNamespaceResponse, error)
}

type platformService struct {
	c    client.Client
	name string
}

func NewPlatformService(name string, c client.Client) PlatformService {
	return &platformService{
		c:    c,
		name: name,
	}
}

func (c *platformService) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...client.CallOption) (*CreateNamespaceResponse, error) {
	req := c.c.NewRequest(c.name, "Platform.CreateNamespace", in)
	out := new(CreateNamespaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformService) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...client.CallOption) (*DeleteNamespaceResponse, error) {
	req := c.c.NewRequest(c.name, "Platform.DeleteNamespace", in)
	out := new(DeleteNamespaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Platform service

type PlatformHandler interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest, *CreateNamespaceResponse) error
	DeleteNamespace(context.Context, *DeleteNamespaceRequest, *DeleteNamespaceResponse) error
}

func RegisterPlatformHandler(s server.Server, hdlr PlatformHandler, opts ...server.HandlerOption) error {
	type platform interface {
		CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, out *CreateNamespaceResponse) error
		DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, out *DeleteNamespaceResponse) error
	}
	type Platform struct {
		platform
	}
	h := &platformHandler{hdlr}
	return s.Handle(s.NewHandler(&Platform{h}, opts...))
}

type platformHandler struct {
	PlatformHandler
}

func (h *platformHandler) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, out *CreateNamespaceResponse) error {
	return h.PlatformHandler.CreateNamespace(ctx, in, out)
}

func (h *platformHandler) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, out *DeleteNamespaceResponse) error {
	return h.PlatformHandler.DeleteNamespace(ctx, in, out)
}
