package middleware

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthMiddleware struct {
	AuthFunc func(ctx context.Context) (bool, error)
}

func (a *AuthMiddleware) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("HERE", metadata.ExtractIncoming(ctx))
	authorized, err := a.AuthFunc(ctx)
	if err != nil {
		return nil, err
	}

	if !authorized {
		return nil, status.Errorf(codes.Unauthenticated, "User is not authorized")
	}

	return handler(ctx, req)
}

func CheckAuth(ctx context.Context) (bool, error) {
	md := metadata.ExtractIncoming(ctx)

	authorizationHeader := md["authorization"]

	if authorizationHeader == nil || authorizationHeader[0] == "" {
		return false, nil
	}

	token := authorizationHeader[0]

	_ = token

	return true, nil
}
