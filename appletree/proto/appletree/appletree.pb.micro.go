// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/appletree/appletree.proto

package go_micro_srv_appletree

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Appletree service

type AppletreeService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Appletree_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Appletree_PingPongService, error)
}

type appletreeService struct {
	c    client.Client
	name string
}

func NewAppletreeService(name string, c client.Client) AppletreeService {
	return &appletreeService{
		c:    c,
		name: name,
	}
}

func (c *appletreeService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Appletree.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appletreeService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Appletree_StreamService, error) {
	req := c.c.NewRequest(c.name, "Appletree.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &appletreeServiceStream{stream}, nil
}

type Appletree_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type appletreeServiceStream struct {
	stream client.Stream
}

func (x *appletreeServiceStream) Close() error {
	return x.stream.Close()
}

func (x *appletreeServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *appletreeServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *appletreeServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *appletreeServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appletreeService) PingPong(ctx context.Context, opts ...client.CallOption) (Appletree_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Appletree.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &appletreeServicePingPong{stream}, nil
}

type Appletree_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type appletreeServicePingPong struct {
	stream client.Stream
}

func (x *appletreeServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *appletreeServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *appletreeServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *appletreeServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *appletreeServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *appletreeServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Appletree service

type AppletreeHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Appletree_StreamStream) error
	PingPong(context.Context, Appletree_PingPongStream) error
}

func RegisterAppletreeHandler(s server.Server, hdlr AppletreeHandler, opts ...server.HandlerOption) error {
	type appletree interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Appletree struct {
		appletree
	}
	h := &appletreeHandler{hdlr}
	return s.Handle(s.NewHandler(&Appletree{h}, opts...))
}

type appletreeHandler struct {
	AppletreeHandler
}

func (h *appletreeHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AppletreeHandler.Call(ctx, in, out)
}

func (h *appletreeHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.AppletreeHandler.Stream(ctx, m, &appletreeStreamStream{stream})
}

type Appletree_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type appletreeStreamStream struct {
	stream server.Stream
}

func (x *appletreeStreamStream) Close() error {
	return x.stream.Close()
}

func (x *appletreeStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *appletreeStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *appletreeStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *appletreeStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *appletreeHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.AppletreeHandler.PingPong(ctx, &appletreePingPongStream{stream})
}

type Appletree_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type appletreePingPongStream struct {
	stream server.Stream
}

func (x *appletreePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *appletreePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *appletreePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *appletreePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *appletreePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *appletreePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}