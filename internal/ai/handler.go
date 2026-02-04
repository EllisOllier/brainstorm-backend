// temporary file for prototyping
package ai

import (
	"encoding/json"
	"net/http"

	"google.golang.org/genai"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AiRequest struct {
	Messages []Message `json:"messages"`
}

func (s *GeminiService) ChatToProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	history = append(history, genai.NewContentFromText("Based on the conversation above, generate a project summary including a name, title, and structured pages.", genai.RoleUser))

	res, err := s.Client.Models.GenerateContent(r.Context(), "gemini-2.0-flash", history, nil)
	if err != nil {
		http.Error(w, "AI Error: "+err.Error(), http.StatusInternalServerError) // Don't use log.Fatal
		return
	}

	reply := res.Text()
	json.NewEncoder(w).Encode(map[string]string{
		"reply": reply,
	})
}
