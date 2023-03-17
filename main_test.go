package chatgpt_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/garydevenay/go-chatgpt-client"
	"github.com/jarcoal/httpmock"
)

func TestSendMessage(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client := chatgpt.NewClient("test-api-key")

	mockResponse := chatgpt.ChatResponse{
		ID:    "testid",
		Model: "gpt-3.5-turbo",
		Choices: []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
			FinishReason string `json:"finish_reason"`
			Index        int    `json:"index"`
		}{
			{
				Message: struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				}{
					Role:    "assistant",
					Content: "The capital of France is Paris.",
				},
				FinishReason: "stop",
				Index:        0,
			},
		},
	}

	mockResponseBody, _ := json.Marshal(mockResponse)

	httpmock.RegisterResponder("POST", "https://api.openai.com/v1/chat/completions",
		httpmock.NewStringResponder(http.StatusOK, string(mockResponseBody)))

	messages := []chatgpt.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "What is the capital of France?"},
	}

	response, err := client.SendMessage("gpt-3.5-turbo", messages)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResponse := "The capital of France is Paris."
	if response != expectedResponse {
		t.Errorf("Expected response: %s, got: %s", expectedResponse, response)
	}
}
