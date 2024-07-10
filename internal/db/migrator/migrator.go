package migrator

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Migrator struct {
	db *sql.DB
}

func New(db *sql.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}

func (m *Migrator) Migrate() {
	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)
	if err := goose.UpContext(context.Background(), m.db, "migrations"); err != nil {
		log.Fatalf("goose %v: %v", "up", err)
	}
}
