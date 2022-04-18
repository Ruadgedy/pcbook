package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

// maximum 1 megabyte
const maxImageSize = 1 << 30

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
	imageStore ImageStore
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
	err := server.laptopStore.Save(laptop)
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

	err := server.laptopStore.Search(
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

func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown,"cannot receive image request"))
	}

	laptopId := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for laptop %s with image type %s",laptopId, imageType)

	laptop, err := server.laptopStore.Find(laptopId)
	if err != nil {
		return logError(status.Errorf(codes.Internal,"cannot find laptop: %v",err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.NotFound,"laptop %s does not exists",laptopId))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for{
		// check context error
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown,"cannot receive chunk data: %v",err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received chunk with size: %d", size)

		imageSize += size
		if imageSize>maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument,"image is too large: %d > %d",imageSize,maxImageSize))
		}

		// write slowly
		//time.Sleep(time.Second)

		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v",err))
		}
	}

	imageId, err := server.imageStore.Save(laptopId, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Unknown,"cannot save image: %v",err))
	}

	resp := &pb.UploadImageResponse{
		Id:   imageId,
		Size: uint32(imageSize),
	}

	if err = stream.SendAndClose(resp); err != nil {
		return logError(status.Errorf(codes.Unknown,"cannot send response: %v",err))
	}

	log.Printf("success saved image with id: %s, size: %d", imageId,imageSize)
	return nil
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	return &LaptopServer{
		laptopStore:laptopStore,
		imageStore:imageStore,
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}
