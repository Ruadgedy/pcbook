package service

import (
	"context"
	"errors"
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.GetId())

	if len(laptop.Id) > 0 {
		// check if it's a valid uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is invalid %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate uuid: %v", err)
		}
		laptop.Id = id.String()
	}

	// save the laptop to store
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save the laptop: %v", err)
	}

	log.Printf("saved laptop wiith id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{Id: laptop.Id}
	return res, nil
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}
