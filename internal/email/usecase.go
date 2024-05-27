package email

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"github.com/eegomez/stori-challenge/internal/file"
	"html/template"
	"log"
	"mime/multipart"
)

type UseCase interface {
	SendReport(ctx context.Context, reportEmail ReportEmail) error
}

func NewUseCaseFactory(cfg *configuration.Config) UseCase {
	return newUseCase(cfg, NewRepositoryFactory(cfg))
}

func newUseCase(cfg *configuration.Config, repository Repository) UseCase {
	return &useCaseImpl{
		config:     cfg,
		repository: repository,
		fileUC:     file.NewUseCaseFactory(cfg),
	}
}

type useCaseImpl struct {
	config     *configuration.Config
	repository Repository
	fileUC     file.UseCase
}

func (uc *useCaseImpl) SendReport(ctx context.Context, reportEmail ReportEmail) error {
	htmlReport := buildHTMLReport(
		reportEmail.TotalBalance, reportEmail.TransactionsByMonth, reportEmail.AverageDebit, reportEmail.AverageCredit)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreatePart(map[string][]string{
		"Content-Type": {"text/html; charset=UTF-8"},
	})
	if err != nil {
		log.Fatal(err)
	}
	templateContent, err := uc.fileUC.GetReportTemplateFile(ctx)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.New("emailTemplate").Parse(templateContent)
	if err != nil {
		return fmt.Errorf("failed to parse template, %v", err)
	}
	tmpl.Execute(part, struct {
		Report template.HTML
	}{
		Report: template.HTML(htmlReport),
	})

	imageContent, err := uc.fileUC.GetStoriLogoFile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	imagePart, err := writer.CreatePart(map[string][]string{
		"Content-Type":              {"image/png"},
		"Content-Transfer-Encoding": {"base64"},
		"Content-ID":                {"<image1>"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	encoder := base64.NewEncoder(base64.StdEncoding, imagePart)
	_, err = encoder.Write(imageContent)
	if err != nil {
		return fmt.Errorf("failed to encode attachment, %v", err)
	}
	defer encoder.Close()

	return uc.repository.SendEmail(ctx, reportEmail.DestinationEmailAddress, writer.Boundary(), string(body.Bytes()))
}

func buildHTMLReport(balance float64, transactionsByMonth []map[string]int, averageDebit float64, averageCredit float64) string {
	var transactionsByMonthHTML string
	for _, v := range transactionsByMonth {
		for key, value := range v {
			transactionsByMonthHTML += fmt.Sprintf("<p><strong>Number of transactions in %s:</strong> %d </p>\n", key, value)
		}
	}
	htmlReport := fmt.Sprintf(`<p><strong>Total balance is:</strong> %f </p>
        <p><strong>Average debit amount:</strong> %f </p>
        <p><strong>Average credit amount:</strong> %f </p>
		%s`, balance, averageDebit, averageCredit, transactionsByMonthHTML)
	return htmlReport
}
