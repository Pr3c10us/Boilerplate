package payment

import (
	"github.com/shopspring/decimal"
	"github.com/stripe/stripe-go/v79"
)

type Customer struct {
	Email     string
	FirstName string
	LastName  string
}

type Plan struct {
	Name        string
	Amount      decimal.Decimal
	Interval    Intervals
	Currency    Currency
	Description string
}

type SubscriptionResult struct {
	ClientSecret   string
	SubscriptionID string
}

type TransactionParams struct {
	Amount   decimal.Decimal
	Email    string
	Currency Currency
}

type Intervals interface {
	Stripe() stripe.PriceRecurringInterval
	Paystack() string
}
type interval struct {
	stripeInterval   stripe.PriceRecurringInterval
	paystackInterval string
}

func (i *interval) Stripe() stripe.PriceRecurringInterval {
	return i.stripeInterval
}
func (i *interval) Paystack() string {
	return i.paystackInterval
}
func NewInterval(stripe stripe.PriceRecurringInterval, paystack string) Intervals {
	return &interval{
		stripeInterval:   stripe,
		paystackInterval: paystack,
	}
}

var (
	Day   = NewInterval(stripe.PriceRecurringIntervalDay, "daily")
	Week  = NewInterval(stripe.PriceRecurringIntervalWeek, "weekly")
	Month = NewInterval(stripe.PriceRecurringIntervalMonth, "monthly")
	Year  = NewInterval(stripe.PriceRecurringIntervalYear, "yearly")
)

type Currency interface {
	Stripe() stripe.Currency
	Paystack() string
}
type currency struct {
	stripeCurrency   stripe.Currency
	paystackCurrency string
}

func (i *currency) Stripe() stripe.Currency {
	return i.stripeCurrency
}
func (i *currency) Paystack() string {
	return i.paystackCurrency
}
func NewCurrency(stripe stripe.Currency, paystack string) Currency {
	return &currency{
		stripeCurrency:   stripe,
		paystackCurrency: paystack,
	}
}

var (
	NGN = NewCurrency(stripe.CurrencyNGN, "NGN")
	USD = NewCurrency(stripe.CurrencyUSD, "USD")
)
