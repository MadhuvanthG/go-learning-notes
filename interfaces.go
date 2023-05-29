package main

// Using interfaces to reuse methods
// Requirements-
// You have different interactors to communicate with different downstream services
// But there's a common pattern to build the URL to communicate to those services
// Example:
// A request to products service might look like- https://data.product.com/getProducts?fields=title,color,availableQuantity&byId=94242417
// A request to service that returns all available payment vendors might look like- http://data.vendor.com/getPaymentVendors?fields=title,org,rating&byAvailable=true
// You can see a pattern of how this URL is constructed - https://<base-url>/<path>?<query-params-for-fields-requested>&<query-params-for-filters>
// How can use interfaces to define a method "buildRequestPayload" that can be reused by different interactors?

type productsServicesInteractor struct {
	baseURL string // data.product.com
	version string // 2.5.0
	fields  productsMetadata
}

type productsMetadata struct {
	id    string
	label string
	color string
}

func (p productsServicesInteractor) GetBaseURL() string {
	return p.baseURL
}

func (p productsServicesInteractor) GetVersion() string {
	return p.version
}

func (p productsServicesInteractor) GetProducts() (statusCode string) {
	_ = buildRequestPayload([]string{"id", "title"}, p)

	// make the request

	return "200"
}

type vendorServicesInteractor struct {
	baseURL string // data.vendor.com
	version string // 1.3.0
	fields  vendorServiceFields
}

type vendorServiceFields struct {
	org     string
	name    string
	ratings []string
}

func (v vendorServicesInteractor) GetBaseURL() string {
	return v.baseURL
}

func (v vendorServicesInteractor) GetVersion() string {
	return v.version
}

func (v vendorServicesInteractor) GetVendors() (statusCode string) {
	_ = buildRequestPayload([]string{"id", "title"}, v)

	// Make the request

	return "200"
}

type requestPayloadGeneratable interface {
	GetBaseURL() string
	GetVersion() string
}

func buildRequestPayload(fieldsRquested []string, r requestPayloadGeneratable) string {
	return r.GetBaseURL() + r.GetVersion() + fieldsRquested[0]
}
