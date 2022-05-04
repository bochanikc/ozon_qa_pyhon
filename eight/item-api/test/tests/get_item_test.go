//go:build integration
// +build integration

package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	item "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc"
	"testing"
)

func TestGetItem(t *testing.T) {
	host := "localhost:50051"
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial() err: %v", err)
	}
	defer conn.Close()
	itemClient := item.NewItemClient(conn)

	t.Run("get new item", func(t *testing.T) {
		title := "test title"
		description := "test description"

		// arrange
		createItemReq := item.CreateItemRequest{
			Title:       title,
			Description: description,
		}
		createItemRes, err := itemClient.CreateItem(ctx, &createItemReq)

		//act
		id := createItemRes.Id
		t.Logf("createItemRes.Id: %v", id)
		getItemReq := item.GetItemRequest{
			Id: id,
		}
		getItemRes, err := itemClient.GetItem(ctx, &getItemReq)

		//assert
		require.NoError(t, err)
		assert.Equal(t, description, getItemRes.Description)
		assert.Equal(t, title, getItemRes.Title)
		t.Logf("id: %v, des: %v, title: %v", id, getItemRes.Description, getItemRes.Title)
	})
}
