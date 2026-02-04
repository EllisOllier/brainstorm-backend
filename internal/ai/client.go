package ai

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiService struct {
	Client *genai.Client
}

func NewGeminiService(ctx context.Context, apiKey string) (*GeminiService, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	return &GeminiService{Client: client}, nil
}

func (s *GeminiService) HandleChat(ctx context.Context, history []*genai.Content, message string) (string, error) {
	chat, err := s.Client.Chats.Create(ctx, "gemini-2.0-flash", nil, history)
	if err != nil {
		return "", fmt.Errorf("failed to create chat session: %w", err)
	}

	res, err := chat.SendMessage(ctx, genai.Part{Text: message})
	if err != nil {
		return "", err
	}
	return res.Text(), nil
}
