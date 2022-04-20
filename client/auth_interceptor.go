package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

// 客户端拦截器：在发出请求前，加入JWT Token.
// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {
	authClient  *AuthClient
	// key is full method name, value represents the method needs authorization or not
	authMethods map[string]bool
	// JWT token
	accessToken string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(
	authClient *AuthClient,
	authMethods map[string]bool,
	refreshDuration time.Duration,
) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: authMethods,
	}

	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(duration time.Duration) error {
	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := duration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()
			if err != nil {
				wait = time.Second
			}else {
				wait = duration
			}
		}
	}()

	return nil
}

func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	interceptor.accessToken = accessToken
	log.Printf("token refreshed: %v",accessToken)

	return nil
}

// Unary returns a client interceptor to authenticate unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s",method)

		// if the method needs authorization, only add token to the context
		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx),method,req,reply,cc,opts...)
		}

		// otherwise go directly
		return invoker(ctx,method,req,reply,cc,opts...)
	}
}

// Stream returns a client interceptor to authenticate stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (stream grpc.ClientStream, err error) {
		log.Printf("--> stream interceptor: %s",method)

		// if the method needs authorization, only add token to the context
		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx),desc,cc,method,opts...)
		}

		// otherwise go directly
		return streamer(ctx,desc,cc,method,opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx,"authorization",interceptor.accessToken)
}