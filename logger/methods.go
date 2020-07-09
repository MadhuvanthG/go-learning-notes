package logger

import (
	"time"
)

// REQUIREMENTS
// 1. Log important event in all the services
// examples-
// payment auth
// payment refund init
// cart products add
// cart abandon
// 2. Abstract logging complexity from the app services
// 3. Logging service should be reusable

// Logger represents a type that's used for logging
type Logger struct {
	LogType string // payment service, auth service etc..
	Retries int
}

// LogMessage is a type that you want to use
type LogMessage struct {
	Method          string   // name of the method that the log is coming from
	RequestHeaders  []string // information on headers of the request event
	ResponsePayload []string
}

type splunkPayload struct {
	method   string // service name + method name
	time     time.Time
	host     string // K8s pod that generated the error
	metadata []string
}

// LogEvent is the method you want to use log different app events
func (l Logger) LogEvent(message LogMessage) (string, error) {
	// 1. Use the incoming message to construct the splunk payload
	payload := l.constructSplunkPayload(message)

	// 2. Send the event to Splunk
	status, err := l.sendSplunkEvent(payload)

	// 3. Handle failures/retries and propogate the status back to the calling service
	return status, err
}

func (l Logger) constructSplunkPayload(message LogMessage) splunkPayload {
	switch l.LogType {
	case "Payment service":
		return splunkPayload{
			time:   time.Now(),
			host:   "",
			method: l.LogType + message.Method,
		}
	case "Auth service":
	}

	return splunkPayload{}
}

func (l Logger) sendSplunkEvent(payload splunkPayload) (string, error) {
	// Send the event to server

	// Retry it for l.retires times

	return "200", nil
}
