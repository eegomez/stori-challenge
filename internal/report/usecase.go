package report

import (
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
)

type UseCase interface {
}

func NewUseCaseFactory(config *configuration.Config) UseCase {
	return newUseCase(config, NewRepositoryFactory(config))
}

func newUseCase(config *configuration.Config, repository Repository) UseCase {
	return &useCaseImpl{
		config:     config,
		repository: repository,
	}
}

type useCaseImpl struct {
	config     *configuration.Config
	repository Repository
}
