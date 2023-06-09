# OpenAI Golang SDK [![Go Reference](https://pkg.go.dev/badge/github.com/gopenai/openai-client.svg)](https://pkg.go.dev/github.com/gopenai/openai-client)

This is an unofficial Golang SDK for the OpenAI API. It provides a simple and easy-to-use way to interact with the OpenAI API using Golang.

## Features
- **Easy-to-use**: The SDK provides a simple and easy-to-use interface for interacting with the OpenAI API, allowing you to quickly integrate AI capabilities into your Golang projects.
- **Support for OpenTelemetry**: The SDK comes with built-in support for OpenTelemetry, allowing you to capture RED metrics and spans for traces out-of-the-box.
- **Fully configurable**: The SDK is fully configurable, allowing you to customize its behavior to suit your specific needs.

## Installation

To install the OpenAI Golang SDK, simply run:

```sh
go get github.com/gopenai/openai-client
```

## Usage

The OpenAI Golang SDK can be easily generated using a `_oas/schema.yaml` file that defines specification for the OpenAI API. To generate the SDK, project uses the [ogen](https://github.com/ogen-go/ogen) tool, which is a code generator for OpenAPI specifications. This allows you to easily update the SDK as the OpenAI API changes. If you change schema and want to generate sdk you can run the following command in makefile:

```sh
make generate
```

To use the OpenAI Golang SDK, you will need to have an API key for the OpenAI API. You can obtain an API key by [creating an account](https://beta.openai.com/signup/) on the OpenAI website.

Once you have your API key, you can create a new OpenAI client and start making requests to the API. Here's an example:

```go
package main

import (
	"context"
	"fmt"

	openai "github.com/gopenai/openai-client"
)

func main() {
	// Replace SERVER_URL with your actual SERVER_URL key
	client, err := openai.NewClient("SERVER_URL")
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Once upon a time,"
	req := &openai.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    openai.ChatCompletionRequestMessageRoleUser,
				Content: prompt,
			},
		},
	}
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the completed text
	fmt.Println(response.Choices[0].Message.Value.Content)
}


```

This example generates a text completion using the OpenAI API, based on a prompt of "Once upon a time,". The resulting completed text is printed to the console.

You can also use bearer token auth. Here's an example:

```go
package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/gopenai/openai-client"
)

func main() {
	// Replace SERVER_URL and API_TOKEN with your actual SERVER_URL and API_TOKEN values
	client, err := openai.NewBearerAuthClient(os.Getenv("SERVER_URL"), os.Getenv("API_TOKEN"))
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Introduce himself please"
	req := &openai.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    openai.ChatCompletionRequestMessageRoleUser,
				Content: prompt,
			},
		},
	}
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the completed text
	fmt.Println(response.Choices[0].Message.Value.Content)
}

```

## TODO:
- [x] Bearer auth
- [ ] Streaming-api support

## Contributing

This SDK is open source, and contributions are welcome! If you have any bug reports, feature requests, or patches, please submit them through the GitHub issue tracker and pull request system.

## Licence

This project is licensed under the Apache License 2.0 - see the [LICENSE](https://github.com/gopenai/openai-client/blob/main/LICENSE) file for details.