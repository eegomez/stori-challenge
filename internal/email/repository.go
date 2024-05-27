package email

import (
	"context"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
)

type Repository interface {
	SendEmail(ctx context.Context, destinationEmailAddress string, boundary string, body string) error
}

func NewRepositoryFactory(cfg *configuration.Config) Repository {
	return newDefaultRepository(cfg)
}
