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

	// 模拟一些重任务
	//time.Sleep(time.Second*6)

	// 判断请求是否取消
	if ctx.Err() == context.Canceled {
		log.Println("request is canceled")
		return nil, status.Errorf(codes.Canceled, "request is canceled")
	}

	// 判断请求是否超时
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return nil, status.Errorf(codes.DeadlineExceeded, "deadline is exceeded")
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

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{Id: laptop.Id}
	return res, nil
}

func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)

	err := server.Store.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{Laptop: laptop}

			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("send laptop with id: %s", laptop.Id)
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}
