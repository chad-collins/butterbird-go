package openAiUtils

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
)

// BuildPrompt sends a message to OpenAI with a given prompt and returns a formatted response string.
func BuildPrompt(client *openai.Client, message, prompt, model string) (string, error) {
	ctx := context.Background() // Create a background context

	// Use the default model if none is provided
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	// Construct the full prompt by appending the user's message to the provided prompt
	fullPrompt := prompt + "\n" + message

	params := openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fullPrompt,
			},
		},
	}

	resp, err := client.CreateChatCompletion(ctx, params)
	if err != nil {
		return "", err
	}

	// Check if we have a valid response
	if len(resp.Choices) > 0 && len(resp.Choices[0].Message.Content) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	// Handle the situation where no valid response is received
	return "", errors.New("no valid response received from OpenAI")
}

func KrobyPrompt() string {
	return "Please assist in rewriting the following message from Kroby. The original author has difficulty typing, often uses incorrect words, and doesn't use punctuation, which can lead to confusion in their meaning. Your task is to correct all grammar, rectify any incorrect word usage, and add proper punctuation. Also, adjust the sentence structure to enhance readability and ensure the intended meaning is clear:\n"
}

func GeneralPrompt() string {
	return "Please respond like Farva from Super Troopers:\n"
}
