package services

import (
	"context"

	"github.com/fandujar/baymax/pkg/providers"
	"github.com/nats-io/nats.go"
	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	NatsClient     *nats.Conn
	OpenAIProvider *providers.OpenAIProvider
}

func NewOpenAIService(p *providers.OpenAIProvider, nc *nats.Conn) *OpenAIService {
	return &OpenAIService{
		NatsClient:     nc,
		OpenAIProvider: p,
	}
}

func (s *OpenAIService) ChatCompletion(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := s.OpenAIProvider.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    "gpt-4o-mini",
			Messages: messages,
		},
	)

	return resp.Choices[0].Message.Content, err

}
