package handler

import (
	"context"
	"sync"
	"time"

	pb "platform-test/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/runtime"
	"github.com/micro/go-micro/v2/store"
	"github.com/micro/go-micro/v2/util/log"
)

// Handler implements the proto interface
type Handler struct {
	sync.RWMutex

	auth     auth.Auth
	store    store.Store
	broker   broker.Broker
	config   config.Config
	runtime  runtime.Runtime
	registry registry.Registry

	Health struct {
		Auth     pb.HealthStatus
		Store    pb.HealthStatus
		Broker   pb.HealthStatus
		Config   pb.HealthStatus
		Runtime  pb.HealthStatus
		Registry pb.HealthStatus
	}
}

// NewHandler returns an initialized handler
func NewHandler(srv micro.Service) *Handler {
	h := Handler{
		auth:     srv.Options().Auth,
		store:    store.DefaultStore,
		broker:   srv.Options().Broker,
		runtime:  runtime.DefaultRuntime,
		registry: srv.Options().Registry,
	}

	// setup the ticker to perform the checks
	// every 15 seconds
	go func() {
		ticker := time.NewTicker(3 * time.Second)

		for {
			<-ticker.C
			h.performChecks()
		}
	}()

	return &h
}

// GetHealth returns the health status for the modules
func (h *Handler) GetHealth(ctx context.Context, req *pb.GetHealthRequest, rsp *pb.GetHealthResponse) error {
	h.Lock()
	defer h.Unlock()

	*rsp = pb.GetHealthResponse{
		Auth:     h.Health.Auth,
		Broker:   h.Health.Broker,
		Config:   h.Health.Config,
		Registry: h.Health.Registry,
		Runtime:  h.Health.Runtime,
		Store:    h.Health.Store,
	}

	return nil
}

// performChecks executes a healthcheck on each module
func (h *Handler) performChecks() {
	log.Infof("Performing healthchecks")

	h.Lock()
	defer h.Unlock()

	h.Health.Auth = h.authHealthStatus()
	h.Health.Store = h.storeHealthStatus()
	h.Health.Broker = h.brokerHealthStatus()
	h.Health.Config = h.configHealthStatus()
	h.Health.Runtime = h.runtimeHealthStatus()
	h.Health.Registry = h.registryHealthStatus()
}

func (h *Handler) authHealthStatus() pb.HealthStatus {
	if h.auth.String() != "service" {
		log.Errorf("Auth Misconfigured: %v", h.auth.String())
		return pb.HealthStatus_UNHEALTHY
	}

	if _, err := h.auth.Generate("foobar"); err != nil {
		log.Errorf("Auth Error: %v", err)
		return pb.HealthStatus_UNHEALTHY
	}

	return pb.HealthStatus_HEALTHY
}

func (h *Handler) storeHealthStatus() pb.HealthStatus {
	if h.store.String() != "service" {
		log.Errorf("Store Misconfigured: %v", h.store.String())
		return pb.HealthStatus_UNHEALTHY
	}

	if _, err := h.store.List(); err != nil {
		log.Errorf("Store Error: %v", err)
		return pb.HealthStatus_UNHEALTHY
	}

	return pb.HealthStatus_HEALTHY
}

func (h *Handler) brokerHealthStatus() pb.HealthStatus {
	if h.broker.String() != "service" {
		log.Errorf("Broker Misconfigured: %v", h.broker.String())
		return pb.HealthStatus_UNHEALTHY
	}

	msg := &broker.Message{}
	if err := h.broker.Publish("platform.test", msg); err != nil {
		log.Errorf("Broker Error: %v", err)
		return pb.HealthStatus_UNHEALTHY
	}

	return pb.HealthStatus_HEALTHY
}

// TODO: implement config healthcheck once config is fully integrated
func (h *Handler) configHealthStatus() pb.HealthStatus {
	return pb.HealthStatus_UNKNOWN
}

func (h *Handler) registryHealthStatus() pb.HealthStatus {
	if h.registry.String() != "service" {
		log.Errorf("Registry Misconfigured: %v", h.registry.String())
		return pb.HealthStatus_UNHEALTHY
	}

	if _, err := h.registry.ListServices(); err != nil {
		log.Errorf("Registry Error: %v", err)
		return pb.HealthStatus_UNHEALTHY
	}

	return pb.HealthStatus_HEALTHY
}

func (h *Handler) runtimeHealthStatus() pb.HealthStatus {
	if h.runtime.String() != "service" {
		log.Errorf("Runtime Misconfigured: %v", h.runtime.String())
		return pb.HealthStatus_UNHEALTHY
	}

	if _, err := h.runtime.List(); err != nil {
		log.Errorf("Runtime Error: %v", err)
		return pb.HealthStatus_UNHEALTHY
	}

	return pb.HealthStatus_HEALTHY
}
