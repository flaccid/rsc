package cac // import "gopkg.in/rightscale/rsc.v3/ca/cac"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v3/dispatch"
	"gopkg.in/rightscale/rsc.v3/rsapi"
)

// Self-Service designer client
type Api struct {
	*rsapi.Api
}

// New returns a client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(h string, a rsapi.Authenticator) *Api {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	return &Api{Api: api}
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	details := dispatch.RequestDetails{
		HttpMethod:            method,
		Host:                  a.Host,
		Url:                   actionUrl,
		Params:                params,
		Payload:               payload,
		FetchLocationResource: a.FetchLocationResource,
	}
	return dispatch.Dispatch(&details, a)
}