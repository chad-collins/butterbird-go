// gpt_client_interface.go
package gpt

// GPTClient defines the interface for interacting with a GPT service.
type GPTClient interface {
	CreateClient(token string) GPTClient
	BuildPrompt(message, prompt, model string) (string, error)
}

type ClientFactory func(token string) GPTClient
