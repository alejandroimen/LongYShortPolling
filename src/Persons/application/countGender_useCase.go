package application

import "github.com/alejandroimen/LongYShortPolling.git/src/Persons/domain/repository"

type CountGender struct {
	repo repository.PersonRepository
}

func NewCountGender(repo repository.PersonRepository) *CountGender {
	return &CountGender{repo: repo}
}

func (cg *CountGender) Run() ([]int, error) {
	return cg.repo.CountGender()
}
