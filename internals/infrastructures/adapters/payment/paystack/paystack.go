package paystack

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/rpip/paystack-go"
)

type PaymentRepositoryPaystack struct {
	PaystackClient *paystack.Client
}

func NewPaymentRepositoryStripe(paystackClient *paystack.Client) payment.Repository {
	return &PaymentRepositoryPaystack{
		PaystackClient: paystackClient,
	}
}
