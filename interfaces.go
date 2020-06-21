package main

type requestPayloadGeneratable interface {
	GetBaseUrl() string
	GetFilters() string
	GetResponseParameters() string
}

type commerceServicesInteractor struct {
	baseURL string
	version string
	fields  commerceServiceFields
}

type commerceServiceFields struct {
	id           string
	label        string
	subscription string
}

func (c commerceServicesInteractor) GetBaseURL() string {
	return "https://..."
}

func (c commerceServicesInteractor) GetFilters() string {
	return "https://..."
}

func (c commerceServicesInteractor) GetResponseParameters() string {
	return "https://..."
}

func requestPayloadGenerator(service requestPayloadGeneratable) string {
	return "Request payload"
}
