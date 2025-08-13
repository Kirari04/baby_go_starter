package app

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var LOG zerolog.Logger
var ENV Env
var DB *gorm.DB

type Env struct {
	Addr    string `env:"ADDR" envDefault:"0.0.0.0:8080"`
	WorkDir string `env:"WORK_DIR" envDefault:"./.data"`

	PublicUrl string `env:"PUBLIC_URL" envDefault:"http://localhost:8080"`

	Datbase string `env:"DATABASE" envDefault:"database.sqlite3"`
}

func Init() {
	// configure logger
	LOG = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().
		Caller().
		Timestamp().
		Logger()

	// load environment variables
	if err := godotenv.Load(); err != nil {
		LOG.Warn().Err(err).Msg("failed to load .env file")
	}

	cfg := Env{}
	if err := env.Parse(&cfg); err != nil {
		LOG.Fatal().Err(err).Msg("failed to parse environment variables")
	}
	ENV = cfg
}
