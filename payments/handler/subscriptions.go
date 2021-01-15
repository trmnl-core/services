package handler

import (
	"context"
	"strings"

	pb "github.com/trmnl-core/services/payments/proto"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	stripe "github.com/stripe/stripe-go/v71"
)

// CreateSubscription via the Stripe API, e.g. "Subscribe John Doe to Notes Gold"
func (h *Provider) CreateSubscription(ctx context.Context, req *pb.CreateSubscriptionRequest, rsp *pb.CreateSubscriptionResponse) error {
	id, err := h.getStripeIDForCustomer(req.CustomerType, req.CustomerId)
	if err != nil {
		return err
	}

	itemParam := &stripe.SubscriptionItemsParams{
		Quantity: stripe.Int64(req.Quantity),
	}
	if len(req.PlanId) > 0 {
		itemParam.Plan = stripe.String(req.PlanId)
	}
	if len(req.PriceId) > 0 {
		itemParam.Price = stripe.String(req.PriceId)
	}
	sub, err := h.client.Subscriptions.New(&stripe.SubscriptionParams{
		Customer: stripe.String(id),
		Items: []*stripe.SubscriptionItemsParams{
			itemParam,
		},
	})
	if err == nil {
		rsp.Subscription = serializeSubscription(sub)
		return nil
	}

	// Handle the error
	switch err.(*stripe.Error).Code {
	case stripe.ErrorCodeParameterInvalidEmpty:
		log.Errorf("Error creating subscription: %v", err)
		return errors.BadRequest(h.name, "missing arguments")
	default:
		return errors.InternalServerError(h.name, "Unexpected stripe error: %v", err)
	}
}

func (h *Provider) ListSubscriptions(ctx context.Context, req *pb.ListSubscriptionsRequest, rsp *pb.ListSubscriptionsResponse) error {
	id, err := h.getStripeIDForCustomer(req.CustomerType, req.CustomerId)
	if err != nil {
		return err
	}
	if len(strings.TrimSpace(id)) == 0 {
		return errors.InternalServerError(h.name, "Stripe id for customer type '%v' and id '%v' is empty", req.CustomerType, req.CustomerId)
	}
	iter := h.client.Subscriptions.List(&stripe.SubscriptionListParams{
		Customer: id,
		Plan:     req.PlanId,
		Price:    req.PriceId,
	})
	if iter.Err() != nil {
		return errors.InternalServerError(h.name, "Unexpected stripe error: %v", iter.Err())
	}

	// Loop through and serialize
	rsp.Subscriptions = make([]*pb.Subscription, 0)
	for {
		if !iter.Next() {
			break
		}

		pm := serializeSubscription(iter.Subscription())
		rsp.Subscriptions = append(rsp.Subscriptions, pm)
	}

	return nil
}

// Update subscription quantity
func (h *Provider) UpdateSubscription(ctx context.Context, req *pb.UpdateSubscriptionRequest, rsp *pb.UpdateSubscriptionResponse) error {
	_, err := h.client.Subscriptions.Update(req.SubscriptionId, &stripe.SubscriptionParams{
		Quantity:          stripe.Int64(req.Quantity),
		ProrationBehavior: stripe.String("always_invoice"),
	})
	if err == nil {
		return nil
	}

	// Handle the error
	switch err.(*stripe.Error).Code {
	case stripe.ErrorCodeParameterInvalidEmpty:
		log.Errorf("Error updating subscription: %v", err)
		return errors.BadRequest(h.name, "missing arguments")
	default:
		return errors.InternalServerError(h.name, "Unexpected stripe error: %v", err)
	}
}

func serializeSubscription(pm *stripe.Subscription) *pb.Subscription {
	rsp := &pb.Subscription{
		Id: pm.ID,
	}

	if pm.Items == nil || len(pm.Items.Data) == 0 {
		return rsp
	}

	plan := pm.Items.Data[0].Plan
	if plan != nil {
		rsp.Plan = serializePlan(plan)
	}
	if plan != nil && plan.Product != nil {
		rsp.Product = serializeProduct(plan.Product)
	}
	rsp.Quantity = pm.Items.Data[0].Quantity

	return rsp
}

func serializePlan(pm *stripe.Plan) *pb.Plan {
	var interval pb.PlanInterval
	switch pm.Interval {
	case stripe.PlanIntervalDay:
		interval = pb.PlanInterval_DAY
	case stripe.PlanIntervalWeek:
		interval = pb.PlanInterval_WEEK
	case stripe.PlanIntervalMonth:
		interval = pb.PlanInterval_MONTH
	case stripe.PlanIntervalYear:
		interval = pb.PlanInterval_YEAR
	}

	return &pb.Plan{
		Id:       pm.ID,
		Name:     pm.Nickname,
		Amount:   pm.Amount,
		Currency: string(pm.Currency),
		Interval: interval,
	}
}

func serializeProduct(prod *stripe.Product) *pb.Product {
	return &pb.Product{
		Id:   prod.ID,
		Name: prod.Name,
	}
}

func (h *Provider) CancelSubscription(ctx context.Context, request *pb.CancelSubscriptionRequest, response *pb.CancelSubscriptionResponse) error {
	_, err := h.client.Subscriptions.Cancel(
		request.SubscriptionId,
		&stripe.SubscriptionCancelParams{Prorate: stripe.Bool(true)},
	)
	if err != nil {
		log.Errorf("Error cancelling subscription %s", err)
		return errors.InternalServerError(h.name+".cancelsubcription", "Error cancelling subscription")
	}
	return nil
}
