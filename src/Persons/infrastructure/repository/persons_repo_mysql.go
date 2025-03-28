package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/LongYShortPolling/src/Persons/domain/entities"
)

type PersonRepoMySQL struct {
	db *sql.DB
}

func NewCreatePersonRepoMySQL(db *sql.DB) *PersonRepoMySQL {
	return &PersonRepoMySQL{db: db}
}

func (r *PersonRepoMySQL) Save(Person entities.User) error {
	query := "INSERT INTO persons (name, email, password) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, Person.Name, Person, Person.Password)
	if err != nil {
		return fmt.Errorf("error insertando User: %w", err)
	}
	return nil
}

func (r *PersonRepoMySQL) FindAll() ([]entities.Person, error) {
	query := "SELECT id, name, age, gender FROM persons"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error buscando los persons: %w", err)
	}
	defer rows.Close()

	var Persons []entities.Person
	for rows.Next() {
		var Person entities.Person
		if err := rows.Scan(&Person.ID, &Person.Name, &Person, &Person.Password); err != nil {
			return nil, err
		}
		Person = append(Person, Person)
	}
	return Persons, nil
}
