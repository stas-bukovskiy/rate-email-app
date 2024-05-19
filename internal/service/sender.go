package service

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"usd-uah-emal-subsriber/internal/config"
	"usd-uah-emal-subsriber/pkg/errs"
)

type SenderService struct {
	from   string
	dialer *gomail.Dialer
}

func NewSenderService(conf config.SMTPConfig) *SenderService {
	d := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: conf.InsecureSkipVerify, ServerName: conf.Host}
	return &SenderService{
		from:   conf.From,
		dialer: d,
	}
}

func (s *SenderService) SendEmail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	if err := s.dialer.DialAndSend(m); err != nil {
		return errs.F("failed to send emails: %w", err)
	}

	return nil
}
