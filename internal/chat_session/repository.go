package chatsession

import "database/sql"

type ChatSessionRepository struct {
	db *sql.DB
}
