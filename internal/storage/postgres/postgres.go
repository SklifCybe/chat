package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(databaseUrl string) (*Storage, error) {
	const op = "internal.storage.postgres.New"

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to open postgres - %w", op, err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("databaseUrl", databaseUrl)
		return nil, fmt.Errorf("%s: failed to ping postgres - %w", op, err)
	}

    fmt.Println("db", db)

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Close() error {
    const op = "internal.storage.postgres.Close"

    if err := s.db.Close(); err != nil {
        return fmt.Errorf("%s: failed to close postgres - %w", op, err)
    }

    return nil
}
