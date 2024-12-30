package app

import (
	"flag"
	"strings"
	"time"

	"github.com/oky31/go-rest-skeleton-nca/internal/config"
)

func LoadConfig() config.Config {
	var cfg config.Config
	flag.IntVar(&cfg.Port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", "", "PostgreSQL DSN")

	flag.IntVar(&cfg.Db.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.Db.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.Db.MaxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.Limiter.Rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enable, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.Smtp.Host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.Smtp.Port, "smtp-post", 2525, "SMTP Port")
	flag.StringVar(&cfg.Smtp.Username, "smtp-username", "5c9724a36310c8", "SMTP Username")
	flag.StringVar(&cfg.Smtp.Password, "smtp-password", "d9c1cfd783f780", "SMTP Password")
	flag.StringVar(&cfg.Smtp.Sender, "smtp-sender", "Greenlight <no-reply@greenlight.oky.com>", "SMTP Password")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(s string) error {
		cfg.Cors.TrustedOrigins = strings.Fields(s)
		return nil
	})

	flag.Parse()

	return cfg
}
