package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/sillkiw/gotube/internal/storage"
	"github.com/sillkiw/gotube/internal/videos"
)

const uniqueViolationCode = "23505"

type Storage struct {
	db *sql.DB
}

func New(postgresSql string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", postgresSql)
	if err != nil {
		return nil, fmt.Errorf("%s: open: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("%s: ping: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Create(v videos.Video) (string, error) {
	const op = "storage.postgres.Create"
	const q = `
		INSERT INTO videos(title, size, status) 
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id string
	err := s.db.QueryRow(q, v.Title, v.Size, v.Status).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == uniqueViolationCode {
			return "", fmt.Errorf("%s: insert: %w", op, storage.ErrTitleExists)
		}
		return "", fmt.Errorf("%s: insert: %w", op, err)
	}
	return id, nil
}
