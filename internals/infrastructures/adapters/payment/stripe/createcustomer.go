package stripe

import (
	"fmt"
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/stripe/stripe-go/v79"
	stripeCustomer "github.com/stripe/stripe-go/v79/customer"
)

func (repo *PaymentRepositoryStripe) CreateCustomer(customer payment.Customer) (string, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(customer.Email),
		Name:  stripe.String(fmt.Sprintf("%v %v", customer.FirstName, customer.LastName)),
	}
	createdCustomer, err := stripeCustomer.New(params)
	if err != nil {
		return "", err
	}

	return createdCustomer.ID, nil
}
