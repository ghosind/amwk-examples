# The Examples of AMWK Framework

This repository contains various examples demonstrating how to use the AMWK framework to build web applications. Each example focuses on a specific aspect of the framework, such as building an HTTP web server, handling query parameters, or creating a handler chain. These examples are designed to help you understand how to use the AMWK framework effectively in your own projects.

## Examples

| Name | Description | Links |
|:----:|:-----------:|:----:|
| Hello World (HTTP Adapter) | A simple example of an HTTP web server. | [Detail](web_helloworld/README.md) \| [Code](web_helloworld/app.go) |
| Query Parameters | An example demonstrating how to handle query parameters in a web application. | [Detail](query_params/README.md) \| [Code](query_params/app.go) |
| Request Headers | An example showing how to access and use request headers in your handlers. | [Detail](request_headers/README.md) \| [Code](request_headers/app.go) |
| Request Body | A simple echo server that reads the request body and sends it back in the response. | [Detail](request_body/README.md) \| [Code](request_body/app.go) |
| Handler Chain | An example showing how to create a handler chain to process requests in a modular way. | [Detail](handler_chain/README.md) \| [Code](handler_chain/app.go) |
| State Management | An example illustrating how to manage state in a web application. | [Detail](state_management/README.md) \| [Code](state_management/app.go) |

## How to Use the Examples

Run the following command in the terminal to install the dependencies for all examples:

```bash
go mod tidy
```

Then, navigate to the directory of the example you want to run and execute the `main.go` file. For example:

```bash
cd web_helloworld
go run main.go
```

You can check the README file of each example for specific instructions on how to test the application and see the expected output. Each example is self-contained and demonstrates a specific feature of the AMWK framework, making it easier for you to learn and apply these concepts in your own projects.

## Contributing

If you have any suggestions for new examples or improvements to existing ones, feel free to open an issue or submit a pull request. We welcome contributions from the community to help make this repository a valuable resource for learning the AMWK framework.

## License

This repository is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more details.
