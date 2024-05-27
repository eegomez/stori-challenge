package file

import (
	"context"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"io"
	"log"
	"os"
)

func newDefaultRepository(cfg *configuration.Config) Repository {
	return &defaultRepository{
		config: cfg,
	}
}

type defaultRepository struct {
	config *configuration.Config
}

func (repo *defaultRepository) GetFile(ctx context.Context, filename string) (io.Reader, error) { //TODO: usar context ctx
	file, err := os.Open(fmt.Sprintf("%s", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return file, nil
}
