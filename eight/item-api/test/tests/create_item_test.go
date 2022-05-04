//go:build integration
// +build integration

package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	item "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateItem(t *testing.T) {
	testCases := []struct {
		name        string
		title       string
		description string
	}{
		{
			"default request should pass",
			"test title",
			"test description",
		},
		{
			"long title should pass",
			"This is most longest title in the world!",
			"test description",
		},
		{
			"multiline description should pass",
			"test title",
			`This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
This is most longest description in the world! OMG!!! LOL! EXPLORE YOUR PASSION!
`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := item.CreateItemRequest{
				Title:       tc.title,
				Description: tc.description,
			}

			res, err := itemClient.CreateItem(ctx, &req)

			require.NoError(t, err)
			assert.NotZero(t, res.Id)
			t.Logf("id: %v", res.Id)
		})
	}
	t.Run("empty description should fail", func(t *testing.T) {
		req := item.CreateItemRequest{
			Title: "test title",
		}

		_, err := itemClient.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
		assert.Contains(t, err.Error(), "empty description")
	})
	t.Run("empty title should fail", func(t *testing.T) {
		req := item.CreateItemRequest{
			Description: "empty description",
		}

		_, err := itemClient.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
		assert.Contains(t, err.Error(), "empty title")
	})
	t.Run("empty request should fail", func(t *testing.T) {
		req := item.CreateItemRequest{}

		_, err := itemClient.CreateItem(ctx, &req)

		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument.String(), status.Code(err).String())
		assert.Contains(t, err.Error(), "empty title")
	})
}
