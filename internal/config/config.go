package config

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Db *pgxpool.Config
}

func New() *Config {
	dbConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Failed to create db config, error: ", err)
	}

	return &Config{
		Db: dbConfig,
	}
}
