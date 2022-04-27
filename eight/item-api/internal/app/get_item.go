package app

import (
	"context"
	"errors"
	"log"

	pb "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := s.DB.SelectItem(ctx, in.GetId())
	if err != nil {
		log.Printf("Internal error: %v\n", err)
		return nil, errors.New("failed to get item")
	}
	if item == nil {
		log.Printf("failed search: %v\n", in.GetId())
		return nil, errors.New("can't find this item")
	}

	return &pb.GetItemResponse{Title: item.Title, Description: item.Description}, nil
}
