package service

import (
	"context"
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/Ruadgedy/pcbook/sample"
	"github.com/Ruadgedy/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
	"time"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{Laptop: laptop}

	resp, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, resp.Id, expectedID)

	// check that the laptop is saved on the serevr
	other, err := laptopStore.Find(resp.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}

func requireSameLaptop(t *testing.T, laptop *pb.Laptop, other *pb.Laptop) {
	// 不能直接比较两个protobuf对象，因为对象当中包含很多特殊的生成字段
	// 需要比较两个protobuf对象，需要忽略special filed
	// One easy way is just serializing the objects to JSON
	json1, err := serializer.ProtobufToJSON(laptop)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(other)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

func newTestLaptopClient(t *testing.T, address string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func startTestLaptopServer(t *testing.T, laptopStore LaptopStore) string {
	laptopServer := NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer(grpc.ConnectionTimeout(time.Second * 10))
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	//grpcServer.Serve(listener)  // block call
	//go grpcServer.Serve(listener)  // non block call
	go func() {
		err := grpcServer.Serve(listener)
		require.NoError(t, err)
	}()

	return listener.Addr().String()
}