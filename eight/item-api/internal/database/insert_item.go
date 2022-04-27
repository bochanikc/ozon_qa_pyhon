package database

import (
	"context"
	"fmt"
	"log"

	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
)

func (db *DB) InsertItem(ctx context.Context, item *models.Item) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO items (title, description)
	VALUES ('%v', '%v') RETURNING id;
`, item.Title, item.Description)
	log.Println(query)
	var id int64
	err := db.QueryRowContext(
		ctx,
		query,
	).Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}
	return id, nil
}
