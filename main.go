package main

func main() {
	demoVariablesAndConstants()

	demoArraysAndSlices()

	demoStructAndPointers()

	// InitializeLoggers()
}

// InitializeLoggers initlaizes all the different loggers application wants
// func InitializeLoggers() {
// 	splunkLoggerConnection := "<token>"
// 	// Logger to be used all functions related to Payment services
// 	paymentServiceLogger := Logger{
// 		logType:                "PaymentService",
// 		splunkLoggerConnection: splunkLoggerConnection,
// 	}

// 	paymentServiceLogger.logEvent(LogEventType{
// 		name:    "authorize payment",
// 		message: []string{"user xxx", "amount xxx"},
// 	})
// }
