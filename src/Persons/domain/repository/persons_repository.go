package repository

import "github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/entities"

type PersonRepository interface {
	Save(user entities.Person) error
	GetRecentPersons() ([]entities.Person, error)
	countGender() ([]entities.Person, error)
}
