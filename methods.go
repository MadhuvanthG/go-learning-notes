package main

// REQUIREMENTS
// 1. Log important event in all the services ex. payment auth, payment refund init, cart products add, cart abandon
// 2. Abstract logging complexity from the app services
// 3. Logging service should be reusable

// Logger groups fields related to sending logs to Splunk
type Logger struct {
	logType                string // server log, downstream request log, client log
	splunkLoggerConnection string // connection string to Splunk server
}

// LogEvent represents information of any event
// that occurs in your application
type LogEvent struct {
	name    string
	message []string
}

type splunkPayload struct {
	eventName string
	logLevel  string // info/debug/error
	metadata  []string
	host      string // name of the Kube pod which generates the event
}

// 1. Expose a method that simplifies logging
func (l *Logger) logEvent(event LogEvent) (statusCode string, err error) {
	// a. Convert LogEvent into a payload
	payload := l.constructSplunkPayload(event)
	statusCode, err = l.sendSplunkEvent(payload)
	return
}

// 2. A method to convert the log object into a payload that Splunk expects
func (l *Logger) constructSplunkPayload(event LogEvent) (payload splunkPayload) {
	switch l.logType {
	case "PaymentService":
		payload = splunkPayload{
			eventName: event.name,
			logLevel:  "Info",
			metadata:  []string{"Test", "Log"},
		}

	case "SubscriptionService":
		payload = splunkPayload{
			eventName: event.name,
			logLevel:  "Info",
			metadata:  []string{"Test", "Log"},
		}
	}

	return
}

// 3. send event to splunk
func (l *Logger) sendSplunkEvent(payload splunkPayload) (string, error) {
	return "200", nil
}

// InitializeLoggers initlaizes all the different loggers application wants
func InitializeLoggers() {
	splunkLoggerConnection := "<token>"
	// Logger to be used all functions related to Payment services
	paymentServiceLogger := Logger{
		logType:                "PaymentService",
		splunkLoggerConnection: splunkLoggerConnection,
	}

	paymentServiceLogger.logEvent(LogEvent{
		name:    "authorize payment",
		message: []string{"user xxx", "amount xxx"},
	})
}
