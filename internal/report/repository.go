package report

import (
	"context"
	"errors"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/file"
	"strconv"
)

var ErrInvalidTransactionID = errors.New("invalid transaction ID")
var ErrInvalidTransactionValue = errors.New("invalid transaction value")

type Repository interface {
	readCsvTransactionFile(ctx context.Context) ([]Transaction, error)
}

func NewRepositoryFactory(config *configuration.Config) Repository {
	return newRepository(config, file.NewUseCaseFactory(config))
}

func newRepository(config *configuration.Config, fileUC file.UseCase) Repository {
	return &defaultRepository{
		config: config,
		fileUC: fileUC,
	}
}

type defaultRepository struct {
	config *configuration.Config
	fileUC file.UseCase
}

func (repo *defaultRepository) readCsvTransactionFile(ctx context.Context) ([]Transaction, error) {
	records, err := repo.fileUC.GetTransactionsFile(ctx)
	if err != nil {
		return nil, err
	}
	var transactions []Transaction
	for i, record := range records {
		if i == 0 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidTransactionID, err)
		}

		date := record[1]
		var transaction float64
		if record[2][0] == '-' {
			transaction, err = strconv.ParseFloat(record[2][1:], 64)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrInvalidTransactionValue, err)
			}
		} else {
			transaction, err = strconv.ParseFloat(record[2][0:], 64)
			if err != nil {
				return nil, fmt.Errorf("%w: %v", ErrInvalidTransactionValue, err)
			}
		}

		if record[2][0] == '-' {
			transaction = -transaction
		}

		t := Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
