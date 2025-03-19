package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	r "tablelink/db/nosql"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthMiddleware(serviceName string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Ambil metadata dari request (Header)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}

		// Validasi X-Link-Service
		serviceHeader := md.Get("x-link-service")
		if len(serviceHeader) == 0 || serviceHeader[0] != "be" {
			return nil, errors.New("invalid X-Link-Service header")
		}

		// Validasi Authorization
		authHeader := md.Get("authorization")
		if len(authHeader) == 0 || !strings.HasPrefix(authHeader[0], "Bearer ") {
			return nil, errors.New("missing or invalid Authorization header")
		}

		authUserName := md.Get("authUserName")
		if len(authHeader) == 0 {
			return nil, errors.New("missing or invalid Authorization header")
		}

		token := strings.TrimPrefix(authHeader[0], "Bearer ")
		if token == "" {
			return nil, errors.New("empty token")
		}

		r.GetKey(fmt.Sprintf("token-username-%s", authUserName))

		return handler(ctx, req)
	}
}
