package report

import (
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
)

type Repository interface {
}

func NewRepositoryFactory(config *configuration.Config) Repository {
	return newRepository(config)
}

func newRepository(config *configuration.Config) Repository {
	return &defaultRepository{}
}

type defaultRepository struct {
}
