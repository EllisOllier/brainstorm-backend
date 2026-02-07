package project

import "database/sql"

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(givenDb *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: givenDb,
	}
}

func (r *ProjectRepository) GetTodoById(project_id int, user_id int) (Project, error) {
	var temp Project
	err := r.db.QueryRow("SELECT title, description FROM projects WHERE id = $1 AND user_id = $2", project_id, user_id).Scan(&temp.Title, &temp.Description)
	if err != nil {
		return temp, err
	}

	return temp, nil
}
