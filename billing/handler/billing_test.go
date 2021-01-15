package handler

import (
	"testing"

	malert "github.com/trmnl-core/services/alert/proto/alert/fakes"
	pb "github.com/trmnl-core/services/billing/proto"
	csproto "github.com/trmnl-core/services/customers/proto"
	mcust "github.com/trmnl-core/services/customers/proto/fakes"
	mt "github.com/trmnl-core/services/internal/test"
	nsproto "github.com/trmnl-core/services/namespaces/proto"
	mns "github.com/trmnl-core/services/namespaces/proto/fakes"
	sproto "github.com/trmnl-core/services/payments/proto"
	mprov "github.com/trmnl-core/services/payments/proto/fakes"
	msub "github.com/trmnl-core/services/subscriptions/proto/fakes"
	uproto "github.com/trmnl-core/services/usage/proto"
	musage "github.com/trmnl-core/services/usage/proto/fakes"
	mstore "github.com/micro/micro/v3/service/store"
	"github.com/micro/micro/v3/service/store/memory"

	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	mstore.DefaultStore = memory.NewStore()
	m.Run()
}

func mockedBilling() *Billing {
	nsSvc := &mns.FakeNamespacesService{}
	nsSvc.ListReturns(&nsproto.ListResponse{
		Namespaces: []*nsproto.Namespace{{Id: "ns1"}},
	}, nil)
	nsSvc.ReadReturns(&nsproto.ReadResponse{
		Namespace: &nsproto.Namespace{Id: "ns1", Owners: []string{"someid"}},
	}, nil)
	pSvc := &mprov.FakeProviderService{}
	pSvc.ListSubscriptionsReturns(&sproto.ListSubscriptionsResponse{
		Subscriptions: []*sproto.Subscription{},
	}, nil)
	uSvc := &musage.FakeUsageService{}
	uSvc.ReadReturns(&uproto.ReadResponse{
		Accounts: []*uproto.Account{
			{
				Namespace: "ns1",
				Services:  4,
				Users:     2,
			},
		},
	}, nil)
	custSvc := &mcust.FakeCustomersService{}
	custSvc.ReadReturns(&csproto.ReadResponse{
		Customer: &csproto.Customer{
			Email: "email@address.com",
		},
	}, nil)

	bs := &Billing{
		ns:   nsSvc,
		ss:   pSvc,
		us:   uSvc,
		subs: &msub.FakeSubscriptionsService{},
		config: &Conf{
			additionalServicesPriceID: "servicesprice",
			additionalUsersPriceID:    "usersprice",
			planID:                    "planid",
			maxIncludedServices:       2,
			report:                    false,
			apiKey:                    "none"},
		cs: custSvc,
		as: &malert.FakeAlertService{},
	}

	return bs
}

func TestNoSubscription(t *testing.T) {
	bs := mockedBilling()
	updates, err := bs.calcUpdate("ns1", false)

	g := NewWithT(t)
	g.Expect(err).To(BeNil())
	g.Expect(updates).To(HaveLen(2))
	g.Expect(updates[0].CustomerID).To(Equal("someid"))
	g.Expect(updates[0].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[0].QuantityFrom).To(Equal(int64(0)))
	g.Expect(updates[0].QuantityTo).To(Equal(int64(1)))
	g.Expect(updates[0].PriceID).To(Equal("usersprice"))

	g.Expect(updates[1].CustomerID).To(Equal("someid"))
	g.Expect(updates[1].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[1].QuantityFrom).To(Equal(int64(0)))
	g.Expect(updates[1].QuantityTo).To(Equal(int64(2)))
	g.Expect(updates[1].PriceID).To(Equal("servicesprice"))

	// checking billing apply
	err = saveUpdate(updates[0])
	g.Expect(err).To(BeNil())
	ctx := mt.ContextWithAccount("micro", "foo")
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	subSvc := bs.subs.(*msub.FakeSubscriptionsService)
	g.Expect(subSvc.UpdateCallCount()).To(Equal(1))
	_, req, _ := subSvc.UpdateArgsForCall(0)
	g.Expect(req.PriceID).To(Equal("usersprice"))
	g.Expect(req.Quantity).To(Equal(int64(1)))
	g.Expect(req.OwnerID).To(Equal("someid"))

	err = saveUpdate(updates[1])
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	g.Expect(subSvc.UpdateCallCount()).To(Equal(2))
	_, req, _ = subSvc.UpdateArgsForCall(1)
	g.Expect(req.PriceID).To(Equal("servicesprice"))
	g.Expect(req.Quantity).To(Equal(int64(2)))
	g.Expect(req.OwnerID).To(Equal("someid"))

}

func TestSubscriptionDecrease(t *testing.T) {
	bs := mockedBilling()
	pSvc := &mprov.FakeProviderService{}
	pSvc.ListSubscriptionsReturns(&sproto.ListSubscriptionsResponse{
		Subscriptions: []*sproto.Subscription{
			{
				Plan: &sproto.Plan{
					Id: "servicesprice",
				},
				Quantity: 7,
			},
			{
				Plan: &sproto.Plan{
					Id: "usersprice",
				},
				Quantity: 5,
			},
		},
	}, nil)
	bs.ss = pSvc

	updates, err := bs.calcUpdate("ns1", false)
	g := NewWithT(t)
	g.Expect(err).To(BeNil())

	g.Expect(updates).To(HaveLen(2))
	g.Expect(updates[0].CustomerID).To(Equal("someid"))
	g.Expect(updates[0].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[0].QuantityFrom).To(Equal(int64(5)))
	g.Expect(updates[0].QuantityTo).To(Equal(int64(1)))
	g.Expect(updates[0].PriceID).To(Equal("usersprice"))

	g.Expect(updates[1].CustomerID).To(Equal("someid"))
	g.Expect(updates[1].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[1].QuantityFrom).To(Equal(int64(7)))
	g.Expect(updates[1].QuantityTo).To(Equal(int64(2)))
	g.Expect(updates[1].PriceID).To(Equal("servicesprice"))

	// checking billing apply
	err = saveUpdate(updates[0])
	g.Expect(err).To(BeNil())
	ctx := mt.ContextWithAccount("micro", "foo")
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	subSvc := bs.subs.(*msub.FakeSubscriptionsService)
	g.Expect(subSvc.UpdateCallCount()).To(Equal(1))
	_, req, _ := subSvc.UpdateArgsForCall(0)
	g.Expect(req.PriceID).To(Equal("usersprice"))
	g.Expect(req.Quantity).To(Equal(int64(1)))
	g.Expect(req.OwnerID).To(Equal("someid"))

	err = saveUpdate(updates[1])
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	g.Expect(subSvc.UpdateCallCount()).To(Equal(2))
	_, req, _ = subSvc.UpdateArgsForCall(1)
	g.Expect(req.PriceID).To(Equal("servicesprice"))
	g.Expect(req.Quantity).To(Equal(int64(2)))
	g.Expect(req.OwnerID).To(Equal("someid"))

}

func TestSubscriptionIncrease(t *testing.T) {
	bs := mockedBilling()
	pSvc := &mprov.FakeProviderService{}
	pSvc.ListSubscriptionsReturns(&sproto.ListSubscriptionsResponse{
		Subscriptions: []*sproto.Subscription{
			{
				Plan: &sproto.Plan{
					Id: "servicesprice",
				},
				Quantity: 1,
			},
			{
				Plan: &sproto.Plan{
					Id: "usersprice",
				},
				Quantity: 1,
			},
		},
	}, nil)
	bs.ss = pSvc

	uSvc := &musage.FakeUsageService{}
	uSvc.ReadReturns(&uproto.ReadResponse{
		Accounts: []*uproto.Account{
			{
				Namespace: "ns1",
				Services:  7,
				Users:     5,
			},
		},
	}, nil)
	bs.us = uSvc

	updates, err := bs.calcUpdate("ns1", false)
	g := NewWithT(t)
	g.Expect(err).To(BeNil())

	g.Expect(updates).To(HaveLen(2))
	g.Expect(updates[0].CustomerID).To(Equal("someid"))
	g.Expect(updates[0].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[0].QuantityFrom).To(Equal(int64(1)))
	g.Expect(updates[0].QuantityTo).To(Equal(int64(4)))
	g.Expect(updates[0].PriceID).To(Equal("usersprice"))

	g.Expect(updates[1].CustomerID).To(Equal("someid"))
	g.Expect(updates[1].CustomerEmail).To(Equal("email@address.com"))
	g.Expect(updates[1].QuantityFrom).To(Equal(int64(1)))
	g.Expect(updates[1].QuantityTo).To(Equal(int64(5)))
	g.Expect(updates[1].PriceID).To(Equal("servicesprice"))

	// checking billing apply
	err = saveUpdate(updates[0])
	g.Expect(err).To(BeNil())
	ctx := mt.ContextWithAccount("micro", "foo")
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	subSvc := bs.subs.(*msub.FakeSubscriptionsService)
	g.Expect(subSvc.UpdateCallCount()).To(Equal(1))
	_, req, _ := subSvc.UpdateArgsForCall(0)
	g.Expect(req.PriceID).To(Equal("usersprice"))
	g.Expect(req.Quantity).To(Equal(int64(4)))
	g.Expect(req.OwnerID).To(Equal("someid"))

	err = saveUpdate(updates[1])
	err = bs.Apply(ctx, &pb.ApplyRequest{
		CustomerID: "someid",
		All:        false,
	}, &pb.ApplyResponse{})
	g.Expect(err).To(BeNil())
	g.Expect(subSvc.UpdateCallCount()).To(Equal(2))
	_, req, _ = subSvc.UpdateArgsForCall(1)
	g.Expect(req.PriceID).To(Equal("servicesprice"))
	g.Expect(req.Quantity).To(Equal(int64(5)))
	g.Expect(req.OwnerID).To(Equal("someid"))
}

func TestNoChange(t *testing.T) {
	bs := mockedBilling()
	pSvc := &mprov.FakeProviderService{}
	pSvc.ListSubscriptionsReturns(&sproto.ListSubscriptionsResponse{
		Subscriptions: []*sproto.Subscription{
			{
				Plan: &sproto.Plan{
					Id: "servicesprice",
				},
				Quantity: 2,
			},
			{
				Plan: &sproto.Plan{
					Id: "usersprice",
				},
				Quantity: 1,
			},
		},
	}, nil)
	bs.ss = pSvc
	updates, err := bs.calcUpdate("ns1", false)
	g := NewWithT(t)
	g.Expect(err).To(BeNil())
	g.Expect(updates).To(BeEmpty())
}
