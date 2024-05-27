package report

import (
	"context"
	"errors"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/file"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Repository_readCsvTransactionFile_Successful(t *testing.T) {
	// Given
	repo, dependencies := buildReportRepository(t)
	ctx := context.Background()
	transactionsFile := MockTransactionsFile()
	expectedTransactions := MockTransactions()

	dependencies.fileUC.EXPECT().GetTransactionsFile(ctx).Return(transactionsFile, nil)

	// When
	transactions, err := repo.readCsvTransactionFile(ctx)

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, transactions)
	assert.Equal(t, expectedTransactions, transactions)
}

func Test_Repository_readCsvTransactionFile_ErrorGettingFile(t *testing.T) {
	// Given
	repo, dependencies := buildReportRepository(t)
	ctx := context.Background()
	someError := errors.New("some-error")

	dependencies.fileUC.EXPECT().GetTransactionsFile(ctx).Return(nil, someError)

	// When
	transactions, err := repo.readCsvTransactionFile(ctx)

	// Then
	assert.Nil(t, transactions)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, someError)
}

func Test_Repository_readCsvTransactionFile_IdValueNotAnInteger(t *testing.T) {
	// Given
	repo, dependencies := buildReportRepository(t)
	ctx := context.Background()
	transactionsFile := MockTransactionsFile()
	transactionsFile[1][0] = "not-an-integer"
	dependencies.fileUC.EXPECT().GetTransactionsFile(ctx).Return(transactionsFile, nil)

	// When
	transactions, err := repo.readCsvTransactionFile(ctx)

	// Then
	assert.NotNil(t, err)
	assert.Nil(t, transactions)
	assert.ErrorIs(t, err, ErrInvalidTransactionID)
}

func Test_Repository_readCsvTransactionFile_PositiveValueNotAnInteger(t *testing.T) {
	// Given
	repo, dependencies := buildReportRepository(t)
	ctx := context.Background()
	transactionsFile := MockTransactionsFile()
	transactionsFile[1][2] = "not-a-float"
	dependencies.fileUC.EXPECT().GetTransactionsFile(ctx).Return(transactionsFile, nil)

	// When
	transactions, err := repo.readCsvTransactionFile(ctx)

	// Then
	assert.NotNil(t, err)
	assert.Nil(t, transactions)
	assert.ErrorIs(t, err, ErrInvalidTransactionValue)
}

func Test_Repository_readCsvTransactionFile_NegativeValueNotAnInteger(t *testing.T) {
	// Given
	repo, dependencies := buildReportRepository(t)
	ctx := context.Background()
	transactionsFile := MockTransactionsFile()
	transactionsFile[1][2] = "-not-a-float"
	dependencies.fileUC.EXPECT().GetTransactionsFile(ctx).Return(transactionsFile, nil)

	// When
	transactions, err := repo.readCsvTransactionFile(ctx)

	// Then
	assert.NotNil(t, err)
	assert.Nil(t, transactions)
	assert.ErrorIs(t, err, ErrInvalidTransactionValue)
}

type repositoryDependencies struct {
	controller *gomock.Controller
	config     *configuration.Config
	fileUC     *file.MockUseCase
}

func buildReportRepository(t *testing.T) (Repository, *repositoryDependencies) {
	d := buildReportRepositoryDependencies(t)
	return newRepository(d.config, d.fileUC), d
}
func buildReportRepositoryDependencies(t *testing.T) *repositoryDependencies {
	controller := gomock.NewController(t)
	return &repositoryDependencies{
		controller: controller,
		config:     configuration.MockDefaultConfig(),
		fileUC:     file.NewMockUseCase(controller),
	}
}
