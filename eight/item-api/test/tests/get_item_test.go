//go:build integration
// +build integration

package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	item "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	t.Run("empty id - failed to get item", func(t *testing.T) {
		//act
		getItemReq := item.GetItemRequest{}
		_, err := itemClient.GetItem(ctx, &getItemReq)

		//assert
		require.Error(t, err)
		assert.Equal(t, codes.Unknown.String(), status.Code(err).String())
		t.Logf("Actual error: %v", status.Code(err).String())
	})
}
