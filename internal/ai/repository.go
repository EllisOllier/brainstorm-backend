package ai

import (
	"database/sql"
	"fmt"
)

type AiRepository struct {
	db *sql.DB
}

func NewAiRepository(givenDb *sql.DB) *AiRepository {
	return &AiRepository{
		db: givenDb,
	}
}

func (r *AiRepository) AddProject(project Project, userId int) (int, error) {
	var entryId int
	err := r.db.QueryRow("INSERT INTO projects (title, description, user_id) VALUES ($1, $2, $3) RETURNING id", project.Title, project.Description, userId).Scan(&entryId)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		return -1, err // -1 error value
	}
	return entryId, nil
}

func (r *AiRepository) AddPage(page Page, projectId int, userId int) (int, error) {
	var entryId int
	err := r.db.QueryRow("INSERT INTO pages (title, content, project_id, user_id) VALUES ($1, $2, $3, $4) RETURNING id", page.Title, page.Content, projectId, userId).Scan(&entryId)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		return -1, err // -1 error value
	}
	return entryId, nil
}
