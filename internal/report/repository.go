package report

import (
	"context"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/file"
	"log"
	"strconv"
)

type Repository interface {
	ReadCsvTransactionFile(ctx context.Context) ([]Transaction, error)
}

func NewRepositoryFactory(config *configuration.Config) Repository {
	return newRepository(config)
}

func newRepository(config *configuration.Config) Repository {
	return &defaultRepository{
		fileUC: file.NewUseCaseFactory(config),
	}
}

type defaultRepository struct {
	fileUC file.UseCase
}

func (repo *defaultRepository) ReadCsvTransactionFile(ctx context.Context) ([]Transaction, error) {
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
			log.Fatalf("Error converting Id: %v", err)
		}

		date := record[1]

		transaction, err := strconv.ParseFloat(record[2][1:], 64)
		if err != nil {
			log.Fatalf("Error converting Transaction: %v", err)
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
