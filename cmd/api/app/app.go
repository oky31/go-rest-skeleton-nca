package app

import (
	"log/slog"
	"sync"

	"github.com/oky31/go-rest-skeleton-nca/internal/config"
	"github.com/oky31/go-rest-skeleton-nca/internal/models"
	"github.com/oky31/go-rest-skeleton-nca/internal/mailer"
)

const Version = "1.0.0"

func New(
	config config.Config,
	logger *slog.Logger,
	models models.Models,
	mailer mailer.Mailer,
) *Application {
	return &Application{
		config: config,
		logger: logger,
		models: models,
		mailer: mailer,
	}
}

type Application struct {
	config config.Config
	logger *slog.Logger
	models models.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}
