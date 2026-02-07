package project

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/EllisOllier/brainstorm-backend/internal/middleware"
)

func (s *ProjectService) GetProjectById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	enc := json.NewEncoder(w)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	row, err := s.projectRepository.GetTodoById(id, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	enc.Encode(row)
}
