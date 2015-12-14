# HTTP Pipe

Package of utilities for intercepting http request and response.

To intercept and alter a request or response body it is important to alter the
content length to match the new body length. Alternatively it is possible to set
the content length to -1.

A request interceptor is useful when altering a request before it reaches your
http handlers.

```
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World"))
})

handler = NewRequestInterceptor(func(r *http.Request) *http.Request {
  // Alter the request here
  return r
})(handler)

http.Handle("/", handler)
```

Intercepting responses from http requests.

```
transport := NewInterceptingTransport(http.DefaultTransport, func(res *http.Response, err error) (*http.Response, error) {
    // Alter the response or error here
    return res, err
})

client := &http.Client{
  Transport: transport,
}

r, _ := http.NewRequest("GET", "http://example.com", nil)

res, err := client.Do(r)
```
