package httpipe

import "net/http"

// ResponseInterceptor is a function that given a response will return a
// response. Be aware that the body may be network connected and you must ensure
// that is is closed and drained.
type ResponseInterceptor func(*http.Response, error) (*http.Response, error)

// RequestInterceptor is a function that given a request, returns a request.
// Feel free to alter the request in any way.
type RequestInterceptor func(*http.Request) *http.Request

// NewRequestInterceptor creates a http handler middleware that applies the
// interceptor function to the body.
func NewRequestInterceptor(fn RequestInterceptor) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, fn(r))
		})
	}
}
