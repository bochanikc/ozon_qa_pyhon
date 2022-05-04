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

	req := item.GetItemRequest{
		Id: 1,
	}

	res, err := itemClient.GetItem(ctx, &req)

	require.NoError(t, err)
	assert.NotZero(t, res.Description)
	assert.NotZero(t, res.GetTitle())
	t.Logf("id: %v, des: %v, title: %v", req.Id, res.Description, res.Title)
}
