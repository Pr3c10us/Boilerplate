package paystack

import (
	"errors"
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/rpip/paystack-go"
)

func (repo *PaymentRepositoryPaystack) SubscribeCustomer(CustomerCode string, PlanCode string) (*payment.SubscriptionResult, error) {
	customer, err := repo.PaystackClient.Customer.Get(CustomerCode)
	if err != nil {
		return nil, err
	}

	params := &paystack.TransactionRequest{
		Email:  customer.Email,
		Amount: 0,
		Plan:   PlanCode,
	}

	createdTransaction, err := repo.PaystackClient.Transaction.Initialize(params)
	if err != nil {
		return nil, err
	}
	if createdTransaction["access_code"] == "" {
		return nil, errors.New("failed to initiate transaction")
	}

	return &payment.SubscriptionResult{
		ClientSecret: createdTransaction["access_code"].(string),
	}, nil
}
