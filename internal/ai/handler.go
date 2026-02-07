// temporary file for prototyping
package ai

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/EllisOllier/brainstorm-backend/internal/middleware"
	"google.golang.org/genai"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AiRequest struct {
	Messages []Message `json:"messages"`
}

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Pages       []Page `json:"pages"`
}

type Page struct {
	PageName string `json:"page_name"`
	Content  string `json:"content"`
}

type ProjectWrapper struct {
	Project Project `json:"project"`
}

func (s *AiService) ChatToProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	var req AiRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}

	var history []*genai.Content
	for _, m := range req.Messages {
		role := m.Role
		if role == "ai" {
			role = string(genai.RoleModel)
		}
		history = append(history, genai.NewContentFromText(m.Content, genai.Role(role)))
	}

	history = append(history, genai.NewContentFromText(`Based on the conversation above, generate a project summary. 
Return ONLY a JSON object following this exact structure:
{
  "project": {
    "title": "string",
    "description": "string",
    "pages": [
      {
        "page_name": "string",
        "content": "string"
      }
    ]
  }
}`, genai.RoleUser))

	res, err := s.Client.Models.GenerateContent(r.Context(), "gemini-2.0-flash", history, nil)
	if err != nil {
		http.Error(w, "AI Error: "+err.Error(), http.StatusInternalServerError) // Don't use log.Fatal
		return
	}

	rawReply := res.Text()

	cleanedJson := strings.ReplaceAll(rawReply, "```json", "")
	cleanedJson = strings.ReplaceAll(cleanedJson, "```", "")
	cleanedJson = strings.TrimSpace(cleanedJson)

	var projectResponse ProjectWrapper
	err = json.Unmarshal([]byte(cleanedJson), &projectResponse)
	if err != nil {
		http.Error(w, "Failed to parse AI JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Change _ to projectId when returning to user after implementation is complete
	_, err = s.aiRepository.AddProject(projectResponse.Project, userId)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projectResponse)
}

// Needs reworking to respond with json
// func (s *AiService) HandleChat(w http.ResponseWriter, r *http.Request) {
// 	chat, err := s.Client.Chats.Create(ctx, "gemini-2.0-flash", nil, history)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create chat session: %w", err)
// 	}

// 	res, err := chat.SendMessage(ctx, genai.Part{Text: message})
// 	if err != nil {
// 		return "", err
// 	}
// 	return res.Text(), nil
// }
