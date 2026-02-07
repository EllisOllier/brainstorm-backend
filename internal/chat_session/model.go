package chatsession

type ChatSession struct {
	ID          int    `json:"id"` // maybe change to chat_session_id
	ChatHistory string `json:"chat_history"`
}

type ChatSessionService struct {
	chatSessionRepository *ChatSessionRepository
}

func NewChatSessionService(givenChatSessionRepository *ChatSessionRepository) *ChatSessionService {
	return &ChatSessionService{
		chatSessionRepository: givenChatSessionRepository,
	}
}
