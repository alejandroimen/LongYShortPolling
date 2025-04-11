package application

import (
	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/repository"
	"github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/entities"
)

type GetPersons struct {
	repo repository.PersonRepository
}

func NewGetPersons(repo repository.PersonRepository) *GetPersons {
	return &GetPersons{repo: repo}
}

func (gu *GetPersons) Run() ([]entities.Person, error) {
	Persons, err := gu.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return Persons, nil
}
