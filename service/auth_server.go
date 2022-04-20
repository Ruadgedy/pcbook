package service

import (
	"context"
	"github.com/Ruadgedy/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication（认证服务器）
type AuthServer struct{
	userStore UserStore // 用户存储
	jwtManager *JWTManager  // JWT管理者
	pb.UnimplementedAuthServiceServer   // GRPC前向兼容
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	err := contextError(ctx)
	if err != nil {
		return nil,err
	}
	user, err := server.userStore.Find(req.Username)
	if err != nil {
		return nil,status.Errorf(codes.Internal,"cannot find user:%v",err)
	}

	if user == nil || !user.IsCorrectPassword(req.Password) {
		return nil,status.Errorf(codes.NotFound,"incorrect username or password: %v",err)
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil,status.Errorf(codes.Internal,"cannot generate token: %v",err)
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res,nil
}