package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
	"testing"
)

func TestSelectItem(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=test password=test sslmode=disable dbname=items")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err: %v", err))
	}
	defer db.Close()
	database := DB{DB: db}

	t.Run("select creates data", testSelectItem(&database))
}

func testSelectItem(db *DB) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		item := models.Item{
			Title:       "test title",
			Description: "test description",
		}

		id, _ := db.InsertItem(ctx, &item)

		testItem, err := db.SelectItem(ctx, id)

		require.NoError(t, err)
		assert.Equal(t, item.Title, testItem.Title)
		assert.Equal(t, item.Description, testItem.Description)
		t.Logf("\nid: %v,\ndes: %v,\ntitle: %v", id, testItem.Description, testItem.Title)
	}
}
