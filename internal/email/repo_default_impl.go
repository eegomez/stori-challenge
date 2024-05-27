package email

import (
	"context"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
	"log"
	"net/smtp"
)

type defaultRepository struct {
	config *configuration.Config
}

func newDefaultRepository(cfg *configuration.Config) Repository {
	return &defaultRepository{
		config: cfg,
	}
}

func (repo *defaultRepository) SendEmail(ctx context.Context, destinationEmailAddress string, boundary string, body string) error { //TODO: usar contexto ctx
	from := repo.config.FromEmail
	password := repo.config.FromEmailPassword
	to := []string{
		destinationEmailAddress,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: Test Email\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/related; boundary=\"" + boundary + "\"\r\n" +
		"\r\n" + body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return err
	}

	log.Println("Email sent successfully!")

	return nil
}
