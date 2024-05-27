package report

import (
	"context"
	"errors"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/email"
)

var ErrParseDate = errors.New("error parsing date")

type UseCase interface {
	createReport(ctx context.Context) (*Report, error)
	SendReport(ctx context.Context, destinationEmailAddress string) error
}

func NewUseCaseFactory(config *configuration.Config) UseCase {
	return newUseCase(config, NewRepositoryFactory(config), email.NewUseCaseFactory(config))
}

func newUseCase(config *configuration.Config, repository Repository, emailUC email.UseCase) UseCase {
	return &useCaseImpl{
		config:     config,
		repository: repository,
		emailUC:    emailUC,
	}
}

type useCaseImpl struct {
	config     *configuration.Config
	repository Repository
	emailUC    email.UseCase
}

func (u useCaseImpl) createReport(ctx context.Context) (*Report, error) {
	transactions, err := u.repository.readCsvTransactionFile(ctx)
	if err != nil {
		return nil, err
	}

	var totalBalance float64
	transactionsByMonth := make(map[int]int)
	var totalDebit, totalCredit float64
	var totalDebitTransactions, totalCreditTransactions int
	var month, day int
	var totalTransactions int

	for _, t := range transactions {
		if t.Transaction == 0 {
			continue
		}
		_, err = fmt.Sscanf(t.Date, "%d/%d", &month, &day)
		if err != nil {
			return nil, ErrParseDate
		}
		totalBalance += t.Transaction
		transactionsByMonth[month] += 1
		totalTransactions += 1
		if t.Transaction > 0 {
			totalCredit += t.Transaction
			totalCreditTransactions += 1
		} else {
			totalDebit += t.Transaction
			totalDebitTransactions += 1
		}
	}
	return &Report{
		TotalBalance:        totalBalance,
		TransactionsByMonth: sortAndReplaceMonths(transactionsByMonth),
		AverageDebit:        totalDebit / float64(totalDebitTransactions),
		AverageCredit:       totalCredit / float64(totalCreditTransactions),
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
	err = u.emailUC.SendReport(ctx, reportEmail)
	if err != nil {
		return err
	}
	return nil
}
