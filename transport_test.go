package httpipe_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamiecuthill/httpipe"
)

func TestNewInterceptingTransportStatusOK(t *testing.T) {
	serv := ServeStatus(http.StatusOK)
	defer serv.Close()

	r, err := http.NewRequest("GET", serv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	i := &testInterceptor{}
	_, err = httpipe.NewInterceptingTransport(http.DefaultTransport, i.Intercept).
		RoundTrip(r)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	if !i.Called {
		t.Error("interceptor must be called")
	}
}

func TestNewInterceptingTransportError(t *testing.T) {
	r, err := http.NewRequest("GET", "http://no.way.this.exists.internal/foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	i := &testInterceptor{}
	_, err = httpipe.NewInterceptingTransport(http.DefaultTransport, i.Intercept).
		RoundTrip(r)
	if err == nil {
		t.Error("expected error")
	}
	if !i.Called {
		t.Error("interceptor must be called")
	}
}

func ServeStatus(status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte("this is the body"))
	}))
}

type testInterceptor struct {
	Called bool
}

func (i *testInterceptor) Intercept(res *http.Response, err error) (*http.Response, error) {
	i.Called = true
	return res, err
}
