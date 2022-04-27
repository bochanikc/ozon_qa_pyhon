package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
	item "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
)

type callOption int

const (
	firstId callOption = iota
	withFatal
	withTimeout
)

type MockDB struct {
	callOption
}

func (m MockDB) InsertItem(ctx context.Context, item *models.Item) (int64, error) {
	if m.callOption == firstId {
		return 1, nil
	} else if m.callOption == withFatal {
		return 0, errors.New("FATAL")
	} else if m.callOption == withTimeout {
		for {
			<-ctx.Done()
			return 0, ctx.Err()
		}
	}
	return 0, nil
}
func (m MockDB) SelectItem(ctx context.Context, id int64) (*models.Item, error) {
	return nil, nil
}

func TestCreateItem(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		s := ItemService{DB: MockDB{callOption: firstId}}
		req := item.CreateItemRequest{
			Title:       "test title",
			Description: "test description",
		}

		res, err := s.CreateItem(ctx, &req)

		require.NoError(t, err)
		assert.NotZero(t, res.Id)
	})

	t.Run("error from database", func(t *testing.T) {
		s := ItemService{DB: MockDB{callOption: withFatal}}
		req := item.CreateItemRequest{
			Title:       "test title",
			Description: "test description",
		}

		_, err := s.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.EqualError(t, err, "failed to create item")
	})

	t.Run("with timeout", func(t *testing.T) {
		s := ItemService{DB: MockDB{callOption: withTimeout}, timeout: time.Nanosecond}
		req := item.CreateItemRequest{
			Title:       "test title",
			Description: "test description",
		}

		_, err := s.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.EqualError(t, err, "failed to create item")
	})

	t.Run("empty request", func(t *testing.T) {
		s := ItemService{}
		req := item.CreateItemRequest{}

		_, err := s.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "empty title")
	})
}
