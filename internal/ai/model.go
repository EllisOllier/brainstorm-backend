package ai

import (
	"context"

	"google.golang.org/genai"
)

type ChatSession struct {
	ID          int    `json:"id"` // maybe change to chat_session_id
	ChatHistory string `json:"chat_history"`
}

type AiService struct {
	aiRepository *AiRepository
	Client       *genai.Client
}

func NewAiService(givenAiRepository *AiRepository, ctx context.Context, apiKey string) (*AiService, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	return &AiService{
		aiRepository: givenAiRepository,
		Client:       client,
	}, nil
}
