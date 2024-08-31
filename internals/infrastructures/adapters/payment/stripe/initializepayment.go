package stripe

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

func (repo *PaymentRepositoryStripe) InitializePayment(params payment.TransactionParams) (string, error) {
	paymentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(params.Amount.IntPart() * 100),
		Currency: stripe.String(string(params.Currency.Stripe())),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	initializedPayment, err := paymentintent.New(paymentParams)
	if err != nil {
		return "", err
	}

	return initializedPayment.ClientSecret, nil
}
