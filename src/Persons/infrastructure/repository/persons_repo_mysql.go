package repository

import (
	"database/sql"
	"fmt"

	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/entities"
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
	query := "SELECT * FROM persons"
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

func (r *PersonRepoMySQL) countGender() ([]int, error) {
	var countMan, countTotal int

	// Contar el total de personas
	queryTotal := "SELECT COUNT(*) FROM persons"
	err := r.db.QueryRow(queryTotal).Scan(&countTotal)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo el total de personas: %w", err)
	}

	// Contar las personas con gender='man'
	queryMan := "SELECT COUNT(*) FROM persons WHERE gender='man'"
	err = r.db.QueryRow(queryMan).Scan(&countMan)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo el total de hombres: %w", err)
	}

	// Calcular la cantidad de mujeres
	countWomen := countTotal - countMan

	return []int{countMan, countWomen}, nil
}
