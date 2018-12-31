package cmd

import (
	"context"

	"cloud.google.com/go/datastore"
	pb "github.com/110y/gae-go-grpc/app/api/proto"
	"github.com/google/uuid"
)

type server struct {
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	u := &user{
		ID:   id.String(),
		Name: req.GetUser().GetName(),
	}
	key := datastore.NameKey("User", u.ID, nil)

	_, err = client.Put(ctx, key, u)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   u.ID,
		Name: u.Name,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	key := datastore.NameKey("User", req.Id, nil)
	u := &user{ID: req.Id}

	err := client.Get(ctx, key, u)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   u.ID,
		Name: u.Name,
	}, nil
}
