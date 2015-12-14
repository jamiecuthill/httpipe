package httpipe

import "net/http"

// NewInterceptingTransport extends http.DefaultTransport and applies the given
// interceptor func to the response body if the http status indicates success.
func NewInterceptingTransport(transport http.RoundTripper, fn ResponseInterceptor) http.RoundTripper {
	return RoundTripperFunc(func(r *http.Request) (res *http.Response, err error) {
		res, err = fn(transport.RoundTrip(r))
		return
	})
}

// RoundTripperFunc is a helper type for quickly creating a http.RoundTripper func
type RoundTripperFunc func(r *http.Request) (res *http.Response, err error)

// RoundTrip implements the http.RoundTripper interface
func (f RoundTripperFunc) RoundTrip(r *http.Request) (res *http.Response, err error) {
	res, err = f(r)
	return
}
