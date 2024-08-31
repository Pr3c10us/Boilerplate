package paystack

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/rpip/paystack-go"
)

func (repo *PaymentRepositoryPaystack) CreatePlan(plan payment.Plan) (string, error) {
	amount, _ := plan.Amount.Float64()
	params := &paystack.Plan{
		Name:        plan.Name,
		Description: plan.Description,
		Interval:    plan.Interval.Paystack(),
		Amount:      float32(amount) * 100,
		Currency:    plan.Currency.Paystack(),
	}

	createdPlan, err := repo.PaystackClient.Plan.Create(params)
	if err != nil {
		return "", err
	}

	return createdPlan.PlanCode, nil
}
