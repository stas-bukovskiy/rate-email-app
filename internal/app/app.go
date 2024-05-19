package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"usd-uah-emal-subsriber/internal/config"
	"usd-uah-emal-subsriber/internal/cron"
	"usd-uah-emal-subsriber/internal/model"
	"usd-uah-emal-subsriber/internal/repository"
	"usd-uah-emal-subsriber/internal/server"
	"usd-uah-emal-subsriber/internal/service"
	"usd-uah-emal-subsriber/pkg/logging"
)

func Run() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log := logging.New(conf.Log.Level, false)

	db, err := gorm.Open(sqlite.Open(conf.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database", "error", err)
	}

	if err = db.AutoMigrate(&model.Subscription{}); err != nil {
		log.Fatal("failed to migrate database", "error", err)
	}

	subRepo := repository.NewSubscriptionRepository(db)

	exchange := service.NewExchangeService(conf.ExchangeAPIConfig)
	sender := service.NewSenderService(conf.SMTPConfig)
	subService := service.NewSubscriptionService(subRepo)
	senderCron := cron.NewSenderCronService(log, sender, exchange, subService)
	go senderCron.Start()

	srv := server.NewServer(log, exchange, subService)
	if err = srv.Run(conf.Server.Host, conf.Server.Port); err != nil {
		log.Fatal("failed to run server", "error", err)
	}
}
