// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/notes/notes.proto

package go_micro_api_distributed

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

// Client API for DistributedNotes service

type DistributedNotesService interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...client.CallOption) (*CreateNoteResponse, error)
	UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...client.CallOption) (*UpdateNoteResponse, error)
	DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...client.CallOption) (*DeleteNoteResponse, error)
	ListNotes(ctx context.Context, in *ListNotesRequest, opts ...client.CallOption) (*ListNotesResponse, error)
}

type distributedNotesService struct {
	c    client.Client
	name string
}

func NewDistributedNotesService(name string, c client.Client) DistributedNotesService {
	return &distributedNotesService{
		c:    c,
		name: name,
	}
}

func (c *distributedNotesService) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...client.CallOption) (*CreateNoteResponse, error) {
	req := c.c.NewRequest(c.name, "DistributedNotes.CreateNote", in)
	out := new(CreateNoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedNotesService) UpdateNote(ctx context.Context, in *UpdateNoteRequest, opts ...client.CallOption) (*UpdateNoteResponse, error) {
	req := c.c.NewRequest(c.name, "DistributedNotes.UpdateNote", in)
	out := new(UpdateNoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedNotesService) DeleteNote(ctx context.Context, in *DeleteNoteRequest, opts ...client.CallOption) (*DeleteNoteResponse, error) {
	req := c.c.NewRequest(c.name, "DistributedNotes.DeleteNote", in)
	out := new(DeleteNoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedNotesService) ListNotes(ctx context.Context, in *ListNotesRequest, opts ...client.CallOption) (*ListNotesResponse, error) {
	req := c.c.NewRequest(c.name, "DistributedNotes.ListNotes", in)
	out := new(ListNotesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistributedNotes service

type DistributedNotesHandler interface {
	CreateNote(context.Context, *CreateNoteRequest, *CreateNoteResponse) error
	UpdateNote(context.Context, *UpdateNoteRequest, *UpdateNoteResponse) error
	DeleteNote(context.Context, *DeleteNoteRequest, *DeleteNoteResponse) error
	ListNotes(context.Context, *ListNotesRequest, *ListNotesResponse) error
}

func RegisterDistributedNotesHandler(s server.Server, hdlr DistributedNotesHandler, opts ...server.HandlerOption) error {
	type distributedNotes interface {
		CreateNote(ctx context.Context, in *CreateNoteRequest, out *CreateNoteResponse) error
		UpdateNote(ctx context.Context, in *UpdateNoteRequest, out *UpdateNoteResponse) error
		DeleteNote(ctx context.Context, in *DeleteNoteRequest, out *DeleteNoteResponse) error
		ListNotes(ctx context.Context, in *ListNotesRequest, out *ListNotesResponse) error
	}
	type DistributedNotes struct {
		distributedNotes
	}
	h := &distributedNotesHandler{hdlr}
	return s.Handle(s.NewHandler(&DistributedNotes{h}, opts...))
}

type distributedNotesHandler struct {
	DistributedNotesHandler
}

func (h *distributedNotesHandler) CreateNote(ctx context.Context, in *CreateNoteRequest, out *CreateNoteResponse) error {
	return h.DistributedNotesHandler.CreateNote(ctx, in, out)
}

func (h *distributedNotesHandler) UpdateNote(ctx context.Context, in *UpdateNoteRequest, out *UpdateNoteResponse) error {
	return h.DistributedNotesHandler.UpdateNote(ctx, in, out)
}

func (h *distributedNotesHandler) DeleteNote(ctx context.Context, in *DeleteNoteRequest, out *DeleteNoteResponse) error {
	return h.DistributedNotesHandler.DeleteNote(ctx, in, out)
}

func (h *distributedNotesHandler) ListNotes(ctx context.Context, in *ListNotesRequest, out *ListNotesResponse) error {
	return h.DistributedNotesHandler.ListNotes(ctx, in, out)
}