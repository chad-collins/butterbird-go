package gpt

import (
	"context"
	"errors"
	"fmt"

	"github.com/chad-collins/butterbird-go/internal/logger"
	"github.com/sashabaranov/go-openai"
)

// OpenAIGPTClient is an implementation of GPTClient using OpenAI.
type OpenAIGPTClient struct {
	Client *openai.Client
}

// CreateClient initializes and returns a new OpenAIGPTClient.
func (client *OpenAIGPTClient) CreateClient(token string) GPTClient {
	newClient := openai.NewClient(token)
	return &OpenAIGPTClient{Client: newClient}
}

func (client *OpenAIGPTClient) BuildPrompt(message, prompt, model string) (string, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	fullPrompt := prompt + "\n" + message

	params := openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: fullPrompt},
		},
	}

	logger.Info("Fetching OpenAI response...")
	resp, err := client.Client.CreateChatCompletion(context.Background(), params)
	if err != nil {
		logger.Warn(err, fmt.Sprintf("Error fetching response from OpenAI: %s", err))
		return "", err
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.Content) == 0 {
		noResponseError := errors.New("no valid response received from OpenAI")
		logger.Warn(noResponseError, "No valid response received from OpenAI")
		return "", noResponseError
	}

	logger.Info("OpenAI response received successfully")
	return resp.Choices[0].Message.Content, nil
}
