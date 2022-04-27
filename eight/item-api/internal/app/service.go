package app

import (
	"time"

	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/database"
	pb "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc"
)

// ItemService ...
type ItemService struct {
	pb.ItemServer
	DB database.Database

	timeout time.Duration
}

// NewItemService ...
func NewItemService(db database.Database) *ItemService {
	return &ItemService{
		DB:      db,
		timeout: 10 * time.Second,
	}
}

// RegisterNewItemService ...
func RegisterNewItemService(s *grpc.Server, db database.Database) {
	pb.RegisterItemServer(s, NewItemService(db))
}
