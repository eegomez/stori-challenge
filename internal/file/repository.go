package file

import (
	"context"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"io"
)

const awsImplementation = "aws"

type Repository interface {
	GetFile(ctx context.Context, filename string) (io.Reader, error)
}

func NewRepositoryFactory(cfg *configuration.Config) Repository {
	if cfg.FileImplementation == awsImplementation {
		return newAwsFileRepositoryFactory(cfg)
	}
	return newDefaultRepository(cfg)
}
