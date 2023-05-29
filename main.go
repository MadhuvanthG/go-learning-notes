package main

import (
	"go-learning-notes/logger"
)

func main() {
	// demoVariablesAndConstants()

	// demoArraysAndSlices()

	// demoStructAndPointers()

	// initializePaymentLogger()

	// DemoMaps()
}

func initializePaymentLogger() {
	paymentLogger := logger.Logger{
		LogType: "payment service",
		Retries: 2,
	}

	paymentLogger.LogEvent(logger.LogMessage{
		Method:          "refund",
		RequestHeaders:  []string{},
		ResponsePayload: []string{},
	})
}
