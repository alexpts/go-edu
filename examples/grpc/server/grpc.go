package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/alexpts/edu-go/examples/grpc/contract"
)

type ServiceA struct {
	pb.UnimplementedAServer // технический метод реализует и методы загрушки дает
}

func (s *ServiceA) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}

	return &pb.UserResponse{
		Id:       1,
		Name:     "alex",
		Lastname: "pts",
		Age:      35,
	}, nil
}
