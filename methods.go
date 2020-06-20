package main

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
}

// 1. Expose a method that simplifies logging
func (l *Logger) logEvent(event LogEvent) (statusCode string, err error) {
	// a. Convert LogEvent into a payload
	payload := l.constructSplunkPayload(event)
	statusCode, err = l.sendSplunkEvent(payload)
	return
}

// 2. A method to convert the log object into a payload that Splunk expects
func (l *Logger) constructSplunkPayload(event LogEvent) splunkPayload {
	return splunkPayload{
		eventName: event.name,
		logLevel:  "Info",
		metadata:  []string{"Test", "Log"},
	}
}

// 3. send event to splunk
func (l *Logger) sendSplunkEvent(payload splunkPayload) (string, error) {
	return "200", nil
}

// InitializeLoggers initlaizes all the different loggers application wants
func InitializeLoggers() {
	splunkLoggerConnection := "<token>"
	paymentServiceLogger := Logger{
		logType:                "PaymentService",
		splunkLoggerConnection: splunkLoggerConnection,
	}

	paymentServiceLogger.logEvent(LogEvent{
		name:    "authorize payment",
		message: []string{"user xxx", "amount xxx"},
	})
}
