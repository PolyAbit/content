package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/PolyAbit/content/internal/models"
	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveDirection(ctx context.Context, code string, name string, exams string, description string) (models.Direction, error) {
	const op = "storage.sqlite.SaveUser"

	fmt.Println(code, name, exams, description)
	stmt, err := s.db.Prepare("INSERT INTO direction(code, name, exams, description) VALUES(?, ?, ?, ?)")
	if err != nil {
		return models.Direction{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, code, name, exams, description)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return models.Direction{}, fmt.Errorf("%s: %w", op, models.ErrDirectionExists)
		}

		return models.Direction{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.Direction{}, nil
}