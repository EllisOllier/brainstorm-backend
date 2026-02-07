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

func (r *ProjectRepository) GetProjectById(project_id int, user_id int) (Project, error) {
	var temp Project
	err := r.db.QueryRow("SELECT id, title, description FROM projects WHERE id = $1 AND user_id = $2", project_id, user_id).Scan(&temp.ID, &temp.Title, &temp.Description)
	if err != nil {
		return temp, err
	}

	return temp, nil
}

func (r *ProjectRepository) GetProjects(userId int) ([]Project, error) {
	projects := []Project{}
	rows, err := r.db.Query("SELECT id, title, description FROM projects WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp Project
		rows.Scan(&temp.ID, &temp.Title, &temp.Description)
		projects = append(projects, temp)
	}

	return projects, nil
}
