package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	AppConfig struct {
		SMTPConfig
		ExchangeAPIConfig
		Database struct {
			DSN string `env:"DATABASE_DSN" env-default:"subscriptions.db"`
		}
		Email struct {
			From string `env:"EMAIL_FROM" env-default:"test@mail.com"`
		}
		Server struct {
			Host string `env:"SERVER_HOST" env-default:"0.0.0.0"`
			Port string `env:"SERVER_PORT" env-default:"8080"`
		}
		Log struct {
			Level string `env:"LOG_LEVEL" env-default:"info"`
		}
	}

	ExchangeAPIConfig struct {
		AccessKey string `env:"API_ACCESS_KEY"`
		BaseCcy   string `env:"API_BASE_CCY" env-default:"USD"`
		TargetCcy string `env:"API_TARGET_CCY" env-default:"UAH"`
	}

	SMTPConfig struct {
		Host               string `env:"SMTP_HOST" env-default:"smtp.gmail.com"`
		Port               int    `env:"SMTP_PORT" env-default:"587"`
		Username           string `env:"SMTP_USERNAME"`
		Password           string `env:"SMTP_PASSWORD"`
		From               string `env:"SMTP_FROM"`
		InsecureSkipVerify bool   `env:"SMTP_INSECURE_SKIP_VERIFY" env-default:"false"`
	}
)

func LoadConfig() (*AppConfig, error) {
	cfg := AppConfig{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	return &cfg, nil
}
