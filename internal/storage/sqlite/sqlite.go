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

func (s *Storage) SaveDirection(ctx context.Context, code string, name string, exams string, description string) error {
	const op = "storage.sqlite.SaveUser"

	stmt, err := s.db.Prepare("INSERT INTO direction(code, name, exams, description) VALUES(?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, code, name, exams, description)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %w", op, models.ErrDirectionExists)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetDirections(ctx context.Context) ([]models.Direction, error) {
	const op = "storage.op.GetDirections"

	rows, err := s.db.Query("SELECT * FROM direction")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var directions []models.Direction

	for rows.Next() {
		var direction models.Direction
		if err := rows.Scan(&direction.Id, &direction.Code, &direction.Name, &direction.Description, &direction.Exams); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		directions = append(directions, direction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return directions, nil
}

func (s *Storage) DeleteDirection(ctx context.Context, directionId int64) error {
	const op = "storage.sqlite.DeleteDirection"

	stmt, err := s.db.Prepare("DELETE FROM direction WHERE id=?")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(directionId)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

