package repository

import "github.com/alejandroimen/LongYShortPolling.git.git/src/Persons/domain/entities"

type PersonRepository interface {
	Save(user entities.Person) error
}
