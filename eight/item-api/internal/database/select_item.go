package database

import (
	"context"
	"time"

	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
)

func (db *DB) SelectItem(ctx context.Context, id int64) (*models.Item, error) {
	query := `
		SELECT title, description
		FROM items
		WHERE id = $1;
`
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var res models.Item
	err := db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&res.Title,
		&res.Description,
	)
	if err != nil {
		return nil, err
	}
	return &models.Item{
		Title:       res.Title,
		Description: res.Description,
	}, nil
}
