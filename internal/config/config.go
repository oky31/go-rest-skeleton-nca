package config

import "time"

type Config struct {
    Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  time.Duration
	}
	Limiter struct {
		Rps    float64
		Burst  int
		Enable bool
	}
	Smtp struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
	Cors struct {
		TrustedOrigins []string
	}
}
