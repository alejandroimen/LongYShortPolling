package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/API_HEXAGONAL/src/users/domain/entities"
)

type UserRepoMySQL struct {
	db *sql.DB
}

func NewCreateUserRepoMySQL(db *sql.DB) *UserRepoMySQL {
	return &UserRepoMySQL{db: db}
}

func (r *UserRepoMySQL) Save(User entities.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, User.Name, User.Email, User.Password)
	if err != nil {
		return fmt.Errorf("error insertando User: %w", err)
	}
	return nil
}
