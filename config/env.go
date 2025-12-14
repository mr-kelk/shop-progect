package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	// --- APP INFO ---
	AppName string `env:"APP_NAME"  envDefault:"shop-progect" `
	AppEnv  string `env:"APP_ENV,required" envDefault:"local"`

	// --- APP SERVER ---
	AppHost    string `env:"APP_HOST" envDefault:"127.0.0.1"`
	AppPort    int    `env:"APP_PORT" envDefault:"8000"`
	SessionKey string `env:"SESSION_KEY,required"`

	// --- ORACLE DATABASE ---
	OracleUser    string `env:"ORACLE_USER,required"`
	OraclePass    string `env:"ORACLE_PASS,required"`
	OracleHost    string `env:"ORACLE_HOST" envDefault:"127.0.0.1"`
	OraclePort    int    `env:"ORACLE_PORT" envDefault:"1521"`
	OracleService string `env:"ORACLE_SERVICE" envDefault:"ORCLPDB1"`

	// --- REDIS DATABASE ---
	RedisHost     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort     int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisUser     string `env:"REDIS_USER" envDefault:""`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`

	// --- EXTRA ---
	DebugSQL bool `env:"DEBUG_SQL" envDefault:"false"`
}

var Cfg *config

func NewConfig() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf(".env file not found: %w", err))
	}

	conf, err := env.ParseAsWithOptions[config](env.Options{
		OnSet: func(tag string, value interface{}, isDefault bool) {
			if isDefault {
				fmt.Printf("WARNING env[%s] using default '%v'\n", tag, value)
			}
		},
	})
	if err != nil {
		panic(fmt.Errorf("could not parse env: %w", err))
	}
	Cfg = &conf
}
