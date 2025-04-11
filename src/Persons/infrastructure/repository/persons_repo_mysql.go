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

func (r *PersonRepoMySQL) Save(Person entities.Person) error {
	query := "INSERT INTO persons (name, age, gender) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, Person.Name, Person.Age, Person.Gender)
	if err != nil {
		return fmt.Errorf("error insertando Person: %w", err)
	}
	return nil
}

func (r *PersonRepoMySQL) GetRecentPersons(lastID int) ([]entities.Person, error) {
	rows, err := r.db.Query("SELECT id, nombre, edad, genero FROM personas WHERE id > ?", lastID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var personas []entities.Person
	for rows.Next() {
		var p entities.Person
		var id int 
		err := rows.Scan(&id, &p.Name, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		personas = append(personas, p)
	}
	return personas, nil
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
