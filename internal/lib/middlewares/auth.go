package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthMiddleware struct {
	AuthFunc func(ctx context.Context) (bool, context.Context, error)
}

type UserClaim struct {
	jwt.RegisteredClaims
	Uid   int64
	Email string
}

type userId string
type isAdmin string

const (
	userKey    userId  = "userId"
	isAdminKey isAdmin = "isAdmin"
)

type PermissionProvider interface {
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

func (a *AuthMiddleware) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	authorized, updatedContext, err := a.AuthFunc(ctx)
	if err != nil {
		return nil, err
	}

	if !authorized {
		return nil, status.Errorf(codes.Unauthenticated, "User is not authorized")
	}

	return handler(updatedContext, req)
}

func New(key string, permProvider PermissionProvider) func(context.Context) (bool, context.Context, error) {
	return func(ctx context.Context) (bool, context.Context, error) {
		md := metadata.ExtractIncoming(ctx)

		authorizationHeader := md["authorization"]

		if authorizationHeader == nil || authorizationHeader[0] == "" {
			return false, ctx, nil
		}

		bearerToken := authorizationHeader[0]
		jwtToken := strings.Split(bearerToken, "Bearer ")[1]

		var userClaim UserClaim

		token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err != nil {
			return false, ctx, fmt.Errorf("failed parse jwt: %w", err)
		}

		if !token.Valid {
			return false, ctx, fmt.Errorf("invalid token: %w", err)
		}

		ctx = context.WithValue(ctx, userKey, userClaim.Uid)

		isAdmin, err := permProvider.IsAdmin(ctx, userClaim.Uid)

		if err != nil {
			return false, ctx, fmt.Errorf("failed to get permission: %w", err)
		}

		ctx = context.WithValue(ctx, isAdminKey, isAdmin)

		return true, ctx, nil
	}
}

func UIDFromContext(ctx context.Context) (int64, bool) {
	uid, ok := ctx.Value(userKey).(int64)
	return uid, ok
}

func IsAdminFromContext(ctx context.Context) (bool, bool) {
	isAdmin, ok := ctx.Value(isAdminKey).(bool)
	return isAdmin, ok
}
