package httpipe_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/jamiecuthill/httpipe"
)

func TestRequestInterceptor(t *testing.T) {
	var intercepted bool
	var r1 = &http.Request{}

	interceptor := httpipe.NewRequestInterceptor(func(r *http.Request) *http.Request {
		intercepted = true
		return r1
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !reflect.DeepEqual(r, r1) {
			t.Fatal("Did not received the intercepted request")
		}
	})

	interceptor(handler).ServeHTTP(httptest.NewRecorder(), &http.Request{})

	if !intercepted {
		t.Fatal("Request was not intercepted")
	}
}
