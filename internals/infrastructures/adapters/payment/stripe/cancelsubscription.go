package stripe

import "github.com/stripe/stripe-go/v79/subscription"

func (repo *PaymentRepositoryStripe) CancelSubscription(subscriptionID string) error {
	_, err := subscription.Cancel(subscriptionID, nil)
	if err != nil {
		return err
	}
	return nil
}
