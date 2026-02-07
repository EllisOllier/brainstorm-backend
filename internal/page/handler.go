package page

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/EllisOllier/brainstorm-backend/internal/middleware"
)

func (s *PageService) GetPageById(w http.ResponseWriter, r *http.Request) {
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

	row, err := s.pageRepository.GetPageById(id, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found: 404", http.StatusNotFound)
			return
		}
		http.Error(w, "Error getting users project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	enc.Encode(row)
}

func (s *PageService) GetPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rawId := r.Context().Value(middleware.UserIdKey)
	userId, ok := rawId.(int)
	if !ok {
		http.Error(w, "Could not find user ID", http.StatusUnauthorized)
		return
	}

	enc := json.NewEncoder(w)

	rows, err := s.pageRepository.GetPages(userId)
	if err != nil {
		http.Error(w, "Error getting users projects", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	enc.Encode(rows)
}
