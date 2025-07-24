package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const metadataAuthorizationKey = "authorization"

func NewAuthTokenPropogator() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			autmd := md.Get(metadataAuthorizationKey)
			if len(autmd) > 0 {
				ctx = metadata.AppendToOutgoingContext(ctx, metadataAuthorizationKey, autmd[0])
			}
		}

		return handler(ctx, req)
	}
}
