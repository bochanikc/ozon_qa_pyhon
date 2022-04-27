package app

import (
	"context"
	"errors"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/models"
	pb "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *ItemService) CreateItem(ctx context.Context, in *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	item, err := convertItem(in)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.DB.InsertItem(ctx, item)
	if err != nil {
		log.Printf("Internal error: %v\n", err)
		return nil, errors.New("failed to create item")
	}

	return &pb.CreateItemResponse{Id: id}, nil
}

func convertItem(in *pb.CreateItemRequest) (*models.Item, error) {
	if in.GetTitle() == "" {
		return nil, errors.New("empty title")
	}
	if in.GetDescription() == "" {
		return nil, errors.New("empty description")
	}
	return &models.Item{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}, nil
}
