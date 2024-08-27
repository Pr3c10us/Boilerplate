package payment

type Repository interface {
	CreateCustomer(customer Customer) (string, error)
	CreatePlan(plan Plan) (string, error)
	SubscribeCustomer(CustomerCode string, PlanCode string) (*SubscriptionResult, error)
	CancelSubscription(SubscriptionID string) error
	UpdateSubscription(SubscriptionID string, PlanCode string) error
	InitializePayment(params TransactionParams) (string, error)
}
