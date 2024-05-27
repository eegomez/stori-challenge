package report

import (
	"context"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/email"
	"log"
)

type UseCase interface {
	createReport(ctx context.Context) (*Report, error)
	SendReport(ctx context.Context, destinationEmailAddress string) error
}

func NewUseCaseFactory(config *configuration.Config) UseCase {
	return newUseCase(config, NewRepositoryFactory(config), email.NewUseCaseFactory(config))
}

func newUseCase(config *configuration.Config, repository Repository, emailUsecase email.UseCase) UseCase {
	return &useCaseImpl{
		config:          config,
		repository:      repository,
		emailRepository: emailUsecase,
	}
}

type useCaseImpl struct {
	config          *configuration.Config
	repository      Repository
	emailRepository email.UseCase
}

func (u useCaseImpl) createReport(ctx context.Context) (*Report, error) {
	transactions, err := u.repository.ReadCsvTransactionFile(ctx)
	if err != nil {
		return nil, err
	}

	var totalBalance float64
	transactionsByMonth := make(map[int]int)
	var totalDebit float64
	var totalCredit float64
	var month, day int
	var totalTransactions int

	for _, t := range transactions {
		if t.Transaction == 0 {
			continue
		}
		_, err = fmt.Sscanf(t.Date, "%d/%d", &month, &day)
		if err != nil {
			log.Println("Error parsing date:", err)
			return nil, err
		}
		totalBalance += t.Transaction
		transactionsByMonth[month] += 1
		totalTransactions += 1
		if t.Transaction > 0 {
			totalCredit += t.Transaction
		} else {
			totalDebit += t.Transaction
		}
	}
	return &Report{
		TotalBalance:        totalBalance,
		TransactionsByMonth: sortAndReplaceMonths(transactionsByMonth),
		AverageDebit:        totalDebit / float64(totalTransactions),
		AverageCredit:       totalCredit / float64(totalTransactions),
	}, nil
}

func (u useCaseImpl) SendReport(ctx context.Context, destinationEmailAddress string) error {
	report, err := u.createReport(ctx)
	if err != nil {
		return err
	}
	reportEmail := email.ReportEmail{
		DestinationEmailAddress: destinationEmailAddress,
		TotalBalance:            report.TotalBalance,
		TransactionsByMonth:     report.TransactionsByMonth,
		AverageDebit:            report.AverageDebit,
		AverageCredit:           report.AverageCredit,
	}
	err = u.emailRepository.SendReport(ctx, reportEmail)
	if err != nil {
		return err
	}
	return nil
}
