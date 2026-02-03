// temporary file for prototyping
package ai

type AI struct {
	Messages string `json:"messages"`
}

type AIService struct {
	ai *AI
}

func NewAIService(givenAi *AI) *AIService {
	return &AIService{
		ai: givenAi,
	}
}
