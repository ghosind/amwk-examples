# Request Body Example

This example illustrates how to read and process the request body in an AMWK web application. This example implements a simple echo server that reads the request body and sends it back in the response.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a POST request with a body to the server. You can use `curl` or any HTTP client to do this:

```bash
curl -X POST -d "Hello, World!" "http://localhost:8000"
# Expected response: "Hello, World!"
curl -X POST -d "Hello, AMWK!" "http://localhost:8000"
# Expected response: "Hello, AMWK!"
```

You should see the response "Hello, World!" when the request body is set to "Hello, World!", and "Hello, AMWK!" when it is set to "Hello, AMWK!".