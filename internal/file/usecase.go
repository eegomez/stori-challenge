package file

import (
	"bytes"
	"context"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"io"
)

const (
	transactionsFile   = "transactions.csv"
	reportTemplateFile = "template.html"
	storiLogoFile      = "stori_logo.png"
)

type UseCase interface {
	GetTransactionsFile(ctx context.Context) ([][]string, error)
	GetReportTemplateFile(ctx context.Context) (string, error)
	GetStoriLogoFile(ctx context.Context) ([]byte, error)
}

func NewUseCaseFactory(cfg *configuration.Config) UseCase {
	return newUseCase(NewRepositoryFactory(cfg))
}

func newUseCase(repository Repository) UseCase {
	return &useCaseImpl{
		repository: repository,
	}
}

type useCaseImpl struct {
	config     *configuration.Config
	repository Repository
}

func (uc *useCaseImpl) GetTransactionsFile(ctx context.Context) ([][]string, error) {
	file, err := uc.repository.GetFile(ctx, transactionsFile)
	if err != nil {
		return nil, err
	}
	return readCsvFile(ctx, file)
}

func (uc *useCaseImpl) GetReportTemplateFile(ctx context.Context) (string, error) {
	file, err := uc.repository.GetFile(ctx, reportTemplateFile)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file content, %v", err)
	}
	return string(body), nil
}

func (uc *useCaseImpl) GetStoriLogoFile(ctx context.Context) ([]byte, error) {
	file, err := uc.repository.GetFile(ctx, storiLogoFile)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content, %v", err)
	}

	return buf.Bytes(), nil
}
