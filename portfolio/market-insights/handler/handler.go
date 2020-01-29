package handler

import (
	"context"

	proto "github.com/micro/services/portfolio/market-insights/proto"
	"github.com/micro/services/portfolio/market-insights/storage"
	// "github.com/micro/go-micro/errors"
)

// New returns an instance of Handler
func New(db storage.Service) *Handler {
	return &Handler{db}
}

// Handler is an object can process RPC requests
type Handler struct {
	db storage.Service
}

// List returns the market insights for a given date
func (h *Handler) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	return nil
}
