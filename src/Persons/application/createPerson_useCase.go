package application

import (
	"fmt"

	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/repository"
	"github.com/alejandroimen/LongYShortPolling/src/Persons/domain/entities"
)

type CreatePerson struct {
	repo repository.PersonRepository
}

func NewCreatePerson(repo repository.PersonRepository) *CreatePerson {
	return &CreatePerson{repo: repo}
}

func (cu *CreatePerson) Run(name string, age int, gender string) error {
	person := entities.Person{Name: name, Age: age, Gender: gender}
	if err := cu.repo.Save(person); err != nil {
		return fmt.Errorf("error al guardar la persona: %w", err)
	}
	return nil
}
