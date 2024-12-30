package main

import (
	"expvar"
	"log/slog"
	"os"
	"runtime"

	_ "github.com/lib/pq"
	"github.com/oky31/go-rest-skeleton-nca/cmd/api/app"
	"github.com/oky31/go-rest-skeleton-nca/internal/db"
	"github.com/oky31/go-rest-skeleton-nca/internal/mailer"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg := app.LoadConfig()

	postgresDb, err := db.OpenConn(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer postgresDb.Close()

	logger.Info("database connection pool established")

	expvar.NewString("version").Set(app.Version)
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))
	expvar.Publish("database", expvar.Func(func() any {
		return postgresDb.Stats()
	}))

	application := app.New(
		cfg,
		logger,
		data.NewModels(postgresDb),
		mailer.New(
			cfg.Smtp.Host,
			cfg.Smtp.Port,
			cfg.Smtp.Username,
			cfg.Smtp.Password,
			cfg.Smtp.Sender),
	)

	if err = application.Server(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
