package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/SklifCybe/chat/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustNew()

	var migrationPath, migrationsTable string

	flag.StringVar(&migrationPath, "migration-path", "", "path to migration")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if migrationPath == "" {
		panic("flag migration-path is required")
	}

	mig, err := migrate.New("file://"+migrationPath, cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	if err := mig.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

    fmt.Println("migrations applied")
}
