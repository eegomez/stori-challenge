package report

import (
	"context"
	"errors"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/email"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Usecase_SendReport_Successful(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	destinationEmailAddress := "some.email.address@gmail.com"
	transactions := MockTransactions()
	reportEmail := email.MockReportEmail()

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(transactions, nil)
	dependencies.emailUC.EXPECT().SendReport(ctx, *reportEmail).Return(nil)

	// When
	err := uc.SendReport(ctx, destinationEmailAddress)

	// Then
	assert.Nil(t, err)
}

func Test_Usecase_SendReport_CreateReportFails(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	destinationEmailAddress := "some.email.address@gmail.com"
	someError := errors.New("some-error")

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(nil, someError)

	// When
	err := uc.SendReport(ctx, destinationEmailAddress)

	// Then
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, someError)
}

func Test_Usecase_SendReport_SendReportFails(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	destinationEmailAddress := "some.email.address@gmail.com"
	transactions := MockTransactions()
	reportEmail := email.MockReportEmail()
	someError := errors.New("some-error")

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(transactions, nil)
	dependencies.emailUC.EXPECT().SendReport(ctx, *reportEmail).Return(someError)

	// When
	err := uc.SendReport(ctx, destinationEmailAddress)

	// Then
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, someError)
}

func Test_Usecase_createReport_Successful(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	transactions := MockTransactions()
	expectedReport := MockReport()

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(transactions, nil)

	// When
	report, err := uc.createReport(ctx)

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, expectedReport, report)
}

func Test_Usecase_createReport_ErrorReadingTransactionsFile(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	someError := errors.New("some-error")

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(nil, someError)

	// When
	report, err := uc.createReport(ctx)

	// Then
	assert.Nil(t, report)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, someError)
}

func Test_Usecase_createReport_ErrorParsingDate(t *testing.T) {
	// Given
	uc, dependencies := buildReportUsecase(t)
	ctx := context.Background()
	transactions := MockTransactions()
	transactions[0].Date = "1/bad-date"

	dependencies.repository.EXPECT().readCsvTransactionFile(ctx).Return(transactions, nil)

	// When
	report, err := uc.createReport(ctx)

	// Then
	assert.Nil(t, report)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrParseDate)
}

type usecaseDependencies struct {
	controller *gomock.Controller
	config     *configuration.Config
	repository *MockRepository
	emailUC    *email.MockUseCase
}

func buildReportUsecase(t *testing.T) (UseCase, *usecaseDependencies) {
	d := buildReportUsecaseDependencies(t)
	return newUseCase(d.config, d.repository, d.emailUC), d
}
func buildReportUsecaseDependencies(t *testing.T) *usecaseDependencies {
	controller := gomock.NewController(t)
	return &usecaseDependencies{
		controller: controller,
		config:     configuration.MockDefaultConfig(),
		repository: NewMockRepository(controller),
		emailUC:    email.NewMockUseCase(controller),
	}
}
