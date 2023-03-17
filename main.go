// Package chatgpt provides a simple Go client for interacting with the OpenAI ChatGPT API.
package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is a ChatGPT API client.
type Client struct {
	APIKey string
}

// NewClient returns a new ChatGPT API client with the provided API key.
func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// Message represents a message object for the ChatGPT API.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request payload for the ChatGPT API.
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// ChatResponse represents a response from the ChatGPT API.
type ChatResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created int64  `json:"created"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

// SendMessage sends a message to the ChatGPT API and returns the assistant's response.
// The model parameter should be either "gpt-3.5-turbo" or "gpt-4".
// The messages parameter should contain an array of Message objects.
func (c *Client) SendMessage(model string, messages []Message) (string, error) {
	if model != "gpt-3.5-turbo" && model != "gpt-4" {
		return "", errors.New("invalid model specified, use 'gpt-3.5-turbo' or 'gpt-4'")
	}

	request := ChatRequest{
		Model:    model,
		Messages: messages,
	}

	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 response from API: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var chatResponse ChatResponse
	if err := json.Unmarshal(respBody, &chatResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return chatResponse.Choices[0].Message.Content, nil
}
