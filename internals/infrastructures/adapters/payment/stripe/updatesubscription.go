package stripe

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/subscription"
)

func (repo *PaymentRepositoryStripe) UpdateSubscription(SubscriptionID string, PlanCode string) error {
	gottenSubscription, err := subscription.Get(SubscriptionID, nil)
	if err != nil {
		return err
	}
	params := &stripe.SubscriptionParams{
		Items: []*stripe.SubscriptionItemsParams{{
			ID:    stripe.String(gottenSubscription.Items.Data[0].ID),
			Price: stripe.String(PlanCode),
		}},
	}

	_, err = subscription.Update(SubscriptionID, params)
	if err != nil {
		return err
	}

	return nil
}
