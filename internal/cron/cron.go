package cron

import (
	"context"
	"fmt"
	"time"

	"usd-uah-emal-subsriber/pkg/errs"
	"usd-uah-emal-subsriber/pkg/logging"
)

type SenderCron struct {
	log        logging.Logger
	sender     SenderService
	exchange   ExchangeService
	subService SubscriptionService
}

func NewSenderCronService(log logging.Logger, sender SenderService, exchange ExchangeService, subService SubscriptionService) *SenderCron {
	return &SenderCron{log: log, sender: sender, exchange: exchange, subService: subService}
}

func (s *SenderCron) Start() {
	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		s.log.Info("sending daily rates")
		if err := s.sendDailyRates(context.Background()); err != nil {
			s.log.Error("failed to send daily rates", "error", err)
		}
		s.log.Info("daily rates sent")
	}
}

func (s *SenderCron) sendDailyRates(ctx context.Context) error {
	subscriptions, err := s.subService.Subscriptions(ctx)
	if err != nil {
		return errs.F("failed to get subscriptions: %w", err)
	}

	rate, err := s.exchange.Rate()
	if err != nil {
		return errs.F("failed to get rate: %v", err)
	}
	subject := "Daily USD to UAH Rate"
	body := fmt.Sprintf("Current USD to UAH rate is: %.2f", rate)

	var emails []string
	for _, sub := range subscriptions {
		emails = append(emails, sub.Email)
	}

	if err = s.sender.SendEmail(emails, subject, body); err != nil {
		return errs.F("failed to send email: %w", err)
	}

	return nil
}
