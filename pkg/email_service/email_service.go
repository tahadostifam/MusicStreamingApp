package email_service

import (
	"github.com/tahadostifam/MusicStreamingApp/config"
	"net/smtp"
)

type EmailService struct {
	smtpConfig config.SMTP
}

func NewEmailService(config *config.Config) *EmailService {
	return &EmailService{config.SMTP}
}

func (s *EmailService) SendMail(receivers []string, message string) error {
	auth := smtp.PlainAuth("", s.smtpConfig.From, s.smtpConfig.Password, s.smtpConfig.Host)

	err := smtp.SendMail(s.smtpConfig.Host+":"+string(rune(s.smtpConfig.Port)), auth, s.smtpConfig.From, receivers, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
