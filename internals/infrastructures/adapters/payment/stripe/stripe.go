package stripe

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/payment"
	"github.com/Pr3c10us/boilerplate/packages/configs"
	"github.com/aws/aws-sdk-go/service/ses"
)

type PaymentRepositoryStripe struct {
}

func NewPaymentRepositoryStripe(environmentVariables *configs.EnvironmentVariables, sesClient *ses.SES) payment.Repository {
	return &PaymentRepositoryStripe{}
}
