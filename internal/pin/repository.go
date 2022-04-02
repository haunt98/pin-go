package pin

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	preparedInsertPin         = "InsertPin"
	preparedSelectPinBySHA256 = "SelectPinBySHA256"

	stmtInitPin = `
CREATE TABLE IF NOT EXISTS pin
(
    pin    TEXT PRIMARY KEY,
    sha256 TEXT
)
`
	stmtSelectPinBySHA256 = `
SELECT pin, sha256
FROM pin
WHERE sha256 = ?;
`
	stmtInsertPin = `
INSERT INTO pin (pin, sha256)
VALUES (?, ?);
`
)

type Repository interface {
	SelectPinBySHA256(ctx context.Context, sha256 string) (Pin, error)
	InsertPin(ctx context.Context, pin Pin) error
}

type repo struct {
	db *sql.DB

	// Prepared statements
	// https://go.dev/doc/database/prepared-statements
	preparedStmts map[string]*sql.Stmt
}

func NewRepository(ctx context.Context, db *sql.DB) (Repository, error) {
	if _, err := db.ExecContext(ctx, stmtInitPin); err != nil {
		return nil, fmt.Errorf("database failed to exec: %w", err)
	}

	var err error
	preparedStmts := make(map[string]*sql.Stmt)
	preparedStmts[preparedInsertPin], err = db.PrepareContext(ctx, stmtInsertPin)
	if err != nil {
		return nil, fmt.Errorf("database failed to prepare context: %w", err)
	}

	preparedStmts[preparedSelectPinBySHA256], err = db.PrepareContext(ctx, stmtSelectPinBySHA256)
	if err != nil {
		return nil, fmt.Errorf("database failed to prepare context: %w", err)
	}

	return &repo{
		db:            db,
		preparedStmts: preparedStmts,
	}, nil
}

func (r *repo) SelectPinBySHA256(ctx context.Context, sha256 string) (Pin, error) {
	var pin Pin

	row := r.preparedStmts[preparedSelectPinBySHA256].QueryRowContext(ctx, sha256)
	if err := row.Scan(
		&pin.Pin,
		&pin.SHA256,
	); err != nil {
		return Pin{}, fmt.Errorf("database failed to scan: %w", err)
	}

	return pin, nil
}

func (r *repo) InsertPin(ctx context.Context, pin Pin) error {
	if _, err := r.preparedStmts[preparedInsertPin].ExecContext(ctx,
		pin.Pin,
		pin.SHA256,
	); err != nil {
		return fmt.Errorf("database failed to exec: %w", err)
	}

	return nil
}
