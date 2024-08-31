package paystack

func (repo *PaymentRepositoryPaystack) CancelSubscription(subscriptionID string) error {
	_, err := repo.PaystackClient.Subscription.Disable(subscriptionID, "")
	return err
}
