package handler

import (
	"context"

	pb "github.com/trmnl-core/services/payments/proto"
	"github.com/micro/micro/v3/service/errors"
	stripe "github.com/stripe/stripe-go/v71"
)

// CreateCustomer via the Stripe API, e.g. "John Doe"
func (h *Provider) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest, rsp *pb.CreateCustomerResponse) error {
	if req.Customer == nil {
		return errors.BadRequest(h.name, "Customer required")
	}
	if req.Customer.Metadata == nil {
		req.Customer.Metadata = make(map[string]string, 0)
	}

	// Check to see if the Customer has already been created
	stripeID, err := h.getStripeIDForCustomer(req.Customer.Type, req.Customer.Id)
	if err != nil {
		return err
	}

	// Construct the params
	var params stripe.CustomerParams
	if email := req.Customer.Metadata["email"]; len(email) > 0 {
		params.Email = stripe.String(email)
	}
	if name := req.Customer.Metadata["name"]; len(name) > 0 {
		params.Name = stripe.String(name)
	}
	if phone := req.Customer.Metadata["phone"]; len(phone) > 0 {
		params.Phone = stripe.String(phone)
	}

	// If the Customer already exists, update using the existing attrbutes
	if len(stripeID) > 0 {
		if _, err := h.client.Customers.Update(stripeID, &params); err != nil {
			return errors.InternalServerError(h.name, "Unexpected stripe update error: %v", err)
		}
		return nil
	}

	// Create the Customer in stripe
	c, err := h.client.Customers.New(&params)
	if err != nil {
		return errors.InternalServerError(h.name, "Unexpected stripe create error: %v", err)
	}

	// Write the ID to the database
	return h.setStripeIDForCustomer(c.ID, req.Customer.Type, req.Customer.Id)
}

func (h *Provider) DeleteCustomer(ctx context.Context, request *pb.DeleteCustomerRequest, response *pb.DeleteCustomerResponse) error {
	if len(request.CustomerId) == 0 {
		return errors.BadRequest(h.name+"deletecustomer", "Customer ID required")
	}
	if len(request.CustomerType) == 0 {
		return errors.BadRequest(h.name+"deletecustomer", "Customer type required")
	}
	sid, err := h.getStripeIDForCustomer(request.CustomerType, request.CustomerId)
	if err != nil {
		return errors.InternalServerError(h.name+"deletecustomer", "Error looking up customer")
	}
	_, err = h.client.Customers.Del(sid, &stripe.CustomerParams{})
	if err != nil {
		return errors.InternalServerError(h.name+"deletecustomer", "Error deleting customer in provider")
	}
	return nil
}

func (h *Provider) GetProviderID(ctx context.Context, request *pb.GetProviderIDRequest, response *pb.GetProviderIDResponse) error {
	if len(request.CustomerId) == 0 {
		return errors.BadRequest(h.name+"getproviderid", "Customer ID required")
	}
	if len(request.CustomerType) == 0 {
		return errors.BadRequest(h.name+"getproviderid", "Customer type required")
	}
	sid, err := h.getStripeIDForCustomer(request.CustomerType, request.CustomerId)
	if err != nil {
		return errors.InternalServerError(h.name+"getproviderid", "Error looking up customer")
	}
	response.ProviderId = sid
	return nil

}
