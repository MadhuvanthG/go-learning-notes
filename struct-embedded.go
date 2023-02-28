package main

import "fmt"

type OrderProcessor struct {
	orderID string
	id      string
	PaymentProcessable
	VisaCardProcessor
	MasterCardProcessor
	retryAttempts int
}

func (op *OrderProcessor) MutateMCID() string {
	op.MasterCardProcessor.MutateMCID()
	return "Only avaiable in MC processor"
}

type VisaCardProcessor struct {
	// fields related to processing VISA payments
	authKey       string
	retryAttempts int

	// we are embedding a logger and metrics publisher
	// so logging and metrics pushing can be done using instances of VisaCardProcessor
	Logger
	MetricsPublisher
}

type MasterCardProcessor struct {
	// fields related to processing Master card payments
	cardConfig string
	id         string
}

type PaymentProcessable interface {
	ProcessPayments() bool
	RetryPayments() bool
	GetPaymentStatus() string
}

func ProcessOrder() {
	// Process order with an associated Visa card payment
	visaProcessor := VisaCardProcessor{
		authKey: "originalAuthKey",
	}
	masterCardProcessor := MasterCardProcessor{
		id: "MC-1",
	}

	// order1Processing := OrderProcessor{
	// 	PaymentProcessable:  &visaProcessor,
	// 	VisaCardProcessor:   visaProcessor,
	// 	MasterCardProcessor: masterCardProcessor,
	// }
	order2Processing := OrderProcessor{
		PaymentProcessable:  &masterCardProcessor,
		VisaCardProcessor:   visaProcessor,
		MasterCardProcessor: masterCardProcessor,
		id:                  "Order-2",
	}

	order2Processing.retryAttempts = 5

	// fmt.Printf("VC authKey before mutation: %s \n", order2Processing.authKey)
	// fmt.Printf("MC ID before mutation: %s \n", order2Processing.MasterCardProcessor.id)
	// order2Processing.MutateAuthKey()
	// fmt.Printf("VC authKey after mutation: %s \n", order2Processing.authKey)
	// order2Processing.MutateMCID()
	// fmt.Printf("MC ID after mutation: %s \n", order2Processing.MasterCardProcessor.id)

	fmt.Printf("Retry attempts before mutation, %d \n", order2Processing.retryAttempts)
	order2Processing.IncreaseRetryPayments()
	fmt.Printf("Retry attempts after mutation, %d \n", order2Processing.retryAttempts)
}

func (v VisaCardProcessor) ProcessPayments() bool {
	// receiver method of the embedded type Logger, is automatically promoted to the outer type "PaymentProcessor",
	// so we could directly access it on the receiver variable "p"
	v.logEvent("starting payment processing")

	// aggregate log messages that corresponds to various steps involved in payment processing
	v.aggregateEvents("initializing connection with payment vendor")
	return true
}

func (v VisaCardProcessor) RetryPayments() bool {
	fmt.Printf("Retry attempt# %d \n", (v.retryAttempts + 1))
	return false
}

func (v VisaCardProcessor) IncreaseRetryPayments() bool {
	fmt.Printf("Retry attempt# %d \n", (v.retryAttempts + 1))
	return false
}

func (v VisaCardProcessor) GetPaymentStatus() string {
	return "Retry in progress"
}

func (v VisaCardProcessor) MutateAuthKey() string {
	v.authKey = "mutated-auth-key"
	return "Retry in progress"
}

func (m *MasterCardProcessor) ProcessPayments() bool {
	return true
}

func (m *MasterCardProcessor) RetryPayments() bool {
	return false
}

func (m *MasterCardProcessor) GetPaymentStatus() string {
	return "Retry in progress"
}

func (m *MasterCardProcessor) MutateMCID() string {
	m.id = "MC-mutated-id"
	return "Only avaiable in MC processor"
}

type Logger struct {
	cid               string
	aggregatedMessage []string
}

func (l Logger) logEvent(message string) {
	fmt.Printf("log related to this cid %s: %s \n", l.cid, message)
}

func (l *Logger) aggregateEvents(message string) {
	fmt.Printf("aggregating message %s related to this cid %s \n", message, l.cid)
	l.aggregatedMessage = append(l.aggregatedMessage, message)
}

type MetricsPublisher struct {
	connectionAddress string
}

func (m MetricsPublisher) pushMetrics() {
	fmt.Printf("sending metrics to server located in the address %s \n", m.connectionAddress)
}
