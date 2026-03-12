# Request Headers Example

This example demonstrates how to access and use request headers in an AMWK web application. Request headers are a fundamental part of HTTP requests, allowing clients to send additional information to the server. The AMWK framework provides an easy way to access these headers in your handlers.

The AMWK framework allows you to access request headers using the `Header`/`HeaderValues`/`Headers` methods on the context. These methods return the value of the specified header, all values for a header, or a map of all headers, respectively.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a request with custom headers to the server. You can use `curl` or any HTTP client to do this:

```bash
curl -H "X-Custom-Header: Hello" "http://localhost:8000"
# Expected response: "Header Value: Hello"
curl -H "X-Custom-Header: World" "http://localhost:8000"
# Expected response: "Header Value: World"
```

You should see the response "Header Value: Hello" when the `X-Custom-Header` is set to "Hello", and "Header Value: World" when it is set to "World".

## API Reference

- `Context.Header(key string) string`: Returns the value of the specified header. If the header is not present, it returns an empty string.
- `Context.HeaderValues(key string) []string`: Returns all values for the specified header. If the header is not present, it returns an empty slice.
- `Context.Headers() http.Header`: Returns a map of all headers, where the keys are the header names and the values are slices of strings containing the header values.
