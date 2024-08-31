package stripe

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
)

type PaymentRepositoryStripe struct {
}

func NewPaymentRepositoryStripe() payment.Repository {
	return &PaymentRepositoryStripe{}
}
