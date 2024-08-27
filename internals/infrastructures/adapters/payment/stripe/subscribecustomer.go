package stripe

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/subscription"
)

func (repo *PaymentRepositoryStripe) SubscribeCustomer(CustomerCode string, PlanCode string) (*payment.SubscriptionResult, error) {
	subscriptionParams := &stripe.SubscriptionParams{
		Customer: stripe.String(CustomerCode),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(PlanCode),
			},
		},
		PaymentBehavior: stripe.String("default_incomplete"),
	}
	subscriptionParams.AddExpand("latest_invoice.payment_intent")
	createdSubscription, err := subscription.New(subscriptionParams)
	if err != nil {
		return nil, err
	}

	return &payment.SubscriptionResult{
		SubscriptionID: createdSubscription.ID,
		ClientSecret:   createdSubscription.LatestInvoice.PaymentIntent.ClientSecret,
	}, nil
}
