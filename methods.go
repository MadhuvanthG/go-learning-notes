package main

import (
	"time"
	_ "time"
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

type logger struct {
	logType string // payment service, auth service etc..
	retries int
}

type logMessage struct {
	method          string   // name of the method that the log is coming from
	requestHeaders  []string // information on headers of the request event
	responsePayload []string
}

type splunkPayload struct {
	method   string // service name + method name
	time     time.Time
	host     string // K8s pod that generated the error
	metadata []string
}

func (l logger) logEventOld(message logMessage) (string, error) {
	// 1. Use the incoming message to construct the splunk payload
	payload := l.constructSplunkPayloadOld(message)

	// 2. Send the event to Splunk
	status, err := l.sendSplunkEventOld(payload)

	// 3. Handle failures/retries and propogate the status back to the calling service
	return status, err
}

func (l logger) constructSplunkPayloadOld(message logMessage) splunkPayload {
	switch message.method {
	case "Payment service":
		return splunkPayload{
			time:   time.Now(),
			host:   "",
			method: l.logType + message.method,
		}
	case "Auth service":
	}

	return splunkPayload{}
}

func (l logger) sendSplunkEventOld(payload splunkPayload) (string, error) {
	// Send the event to server

	// Retry it for l.retires times

	return "200", nil
}
