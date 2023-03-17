# ChatGPT

A simple, developer-friendly Go package for interacting with the OpenAI ChatGPT API.

## Features

- Easy-to-use API client for sending messages to ChatGPT
- Supports both "gpt-3.5-turbo" and "gpt-4" models
- Clear documentation and error messages

## Installation

```bash
go get -u github.com/garydevenay/go-chatgpt-client
```

## Usage
First, import the package:

```go
import "github.com/garydevenay/go-chatgpt-client"
```
Then, create a new ChatGPT client with your API key:

```go
client := chatgpt.NewClient("your-api-key")
```
Now you can send messages to the ChatGPT API:

```go
messages := []chatgpt.Message{
    {Role: "system", Content: "You are a helpful assistant."},
    {Role: "user", Content: "What is the capital of France?"},
}

response, err := client.SendMessage("gpt-3.5-turbo", messages)
if err != nil {
    log.Fatal(err)
}

fmt.Println(response)
```
You can use the same client to send more messages and have an interactive conversation with ChatGPT.

## Documentation
For detailed information about the ChatGPT API, structs, and functions, see the [Godoc Documentation](https://pkg.go.dev/github.com/garydevenay/go-chatgpt-client).

License
MIT