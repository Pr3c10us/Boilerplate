package paystack

import (
	"errors"
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/rpip/paystack-go"
)

func (repo *PaymentRepositoryPaystack) InitializePayment(params payment.TransactionParams) (string, error) {
	amount, _ := params.Amount.Float64()
	paymentParams := &paystack.TransactionRequest{
		Email:    params.Email,
		Amount:   float32(amount) * 100,
		Currency: params.Currency.Paystack(),
	}

	createdTransaction, err := repo.PaystackClient.Transaction.Initialize(paymentParams)
	if err != nil {
		return "", err
	}
	if createdTransaction["access_code"] == "" {
		return "", errors.New("failed to initiate transaction")
	}

	return createdTransaction["access_code"].(string), nil
}
