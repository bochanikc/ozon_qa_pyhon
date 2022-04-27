package database

import (
	"context"
	"fmt"
	"testing"

	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
)

func TestMain(m *testing.M) {
	const (
		dropDB          = `DROP DATABASE IF EXISTS items;`
		createDB        = `CREATE DATABASE items;`
		createItemTable = `CREATE TABLE items
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(40) NOT NULL,
    description TEXT        NOT NULL
);`
	)

	psql, err := sql.Open("postgres", "host=localhost port=5432 user=test password=test sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err: %v", err))
	}
	defer psql.Close()

	_, err = psql.Exec(dropDB)
	if err != nil {
		panic(fmt.Errorf("sql.Exec() err: %v", err))
	}

	_, err = psql.Exec(createDB)
	if err != nil {
		panic(fmt.Errorf("sql.Exec() err: %v", err))
	}

	// teardown
	defer func() {
		_, err = psql.Exec(dropDB)
		if err != nil {
			panic(fmt.Errorf("sql.Exec() err: %v", err))
		}
	}()

	db, err := sql.Open("postgres", "host=localhost port=5432 user=test password=test sslmode=disable dbname=items")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err: %v", err))
	}
	defer db.Close()
	_, err = db.Exec(createItemTable)
	if err != nil {
		panic(fmt.Errorf("sql.Exec() err: %v", err))
	}

	m.Run()
}

func TestInsertItem(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=test password=test sslmode=disable dbname=items")
	if err != nil {
		panic(fmt.Errorf("sql.Open() err: %v", err))
	}
	defer db.Close()
	database := DB{DB: db}

	t.Run("default values should return not zero id", testInsertItem(&database))
}

func testInsertItem(db *DB) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.Background()
		item := models.Item{
			Title:       "test title",
			Description: "test description",
		}

		id, err := db.InsertItem(ctx, &item)

		require.NoError(t, err)
		assert.NotZero(t, id)
	}
}
