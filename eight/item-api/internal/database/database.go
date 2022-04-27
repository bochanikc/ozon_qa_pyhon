package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
)

type Database interface {
	InsertItem(context.Context, *models.Item) (int64, error)
	SelectItem(context.Context, int64) (*models.Item, error)
}

type DB struct {
	*sql.DB
}

func NewDatabase(conn string) (*DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
