package paystack

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/rpip/paystack-go"
)

func (repo *PaymentRepositoryPaystack) CreateCustomer(customer payment.Customer) (string, error) {
	params := &paystack.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}
	createdCustomer, err := repo.PaystackClient.Customer.Create(params)
	if err != nil {
		return "", err
	}
	return createdCustomer.CustomerCode, nil
}
