package interceptor

import (
	"context"

	"github.com/110y/gae-go-grpc/internal/lib/appengine/internal"
	"google.golang.org/grpc"
)

func AppEngineNamespacingUnaryServerInterceptor(namespace string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if namespace != "" {
			ctx = internal.WithAppEngineNamespace(ctx, namespace)
		}
		return handler(ctx, req)
	}
}
