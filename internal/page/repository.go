package page

import "database/sql"

type PageRepository struct {
	db *sql.DB
}

func NewPageRepository(givenDb *sql.DB) *PageRepository {
	return &PageRepository{
		db: givenDb,
	}
}

func (r *PageRepository) GetPageById(page_id int, user_id int) (Page, error) {
	var temp Page
	err := r.db.QueryRow("SELECT id, title, content FROM pages WHERE id = $1 AND user_id = $2", page_id, user_id).Scan(&temp.ID, &temp.Title, &temp.Content)
	if err != nil {
		return temp, err
	}

	return temp, nil
}

func (r *PageRepository) GetPages(userId int) ([]Page, error) {
	projects := []Page{}
	rows, err := r.db.Query("SELECT id, title, content FROM pages WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp Page
		rows.Scan(&temp.ID, &temp.Title, &temp.Content)
		projects = append(projects, temp)
	}

	return projects, nil
}
