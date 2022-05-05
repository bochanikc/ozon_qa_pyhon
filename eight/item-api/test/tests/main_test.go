//go:build integration
// +build integration

package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	item "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/test/internal/config"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/test/internal/interceptor"
	"google.golang.org/grpc"
)

var (
	ctx        context.Context
	itemClient item.ItemClient
)

func TestMain(m *testing.M) {
	// setup integration tests
	cfg := config.ProcessConfig()
	ctx = context.Background()

	conn, err := grpc.Dial(
		cfg.Host,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Timeout(time.Second)),
		grpc.WithUnaryInterceptor(interceptor.Duration()),
	)
	if err != nil {
		panic(fmt.Errorf("grpc.Dial() err: %v", err))
	}
	// teardown integration tests "defer"
	defer conn.Close()

	itemClient = item.NewItemClient(conn)

	m.Run()
}
