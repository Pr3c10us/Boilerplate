package stripe

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
)

func (repo *PaymentRepositoryStripe) CreatePlan(plan payment.Plan) (string, error) {
	productParams := &stripe.ProductParams{
		Name:        stripe.String(plan.Name),
		Description: stripe.String(plan.Description),
	}
	createdProduct, err := product.New(productParams)
	if err != nil {
		return "", err
	}

	priceParams := &stripe.PriceParams{
		Currency: stripe.String(string(plan.Currency.Stripe())), // change to support required currency
		Product:  stripe.String(createdProduct.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(plan.Interval.Stripe())),
		},
		UnitAmount: stripe.Int64(plan.Amount.IntPart()),
	}
	createdPrice, err := price.New(priceParams)
	if err != nil {
		return "", err
	}

	return createdPrice.ID, nil
}
