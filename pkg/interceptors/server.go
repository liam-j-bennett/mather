package interceptors

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	apiKeyHeaderKey = "api_key"
)

// TODO: Change apiKey into a ...Options format instantiation
func New(log *logrus.Logger, apiKey string) grpc.ServerOption {
	interceptor := unaryServerInterceptor(log, apiKey)
	return grpc.UnaryInterceptor(interceptor)
}

func unaryServerInterceptor(log *logrus.Logger, apiKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		start := time.Now()
		defer log.Infoln("Endpoint:", info.FullMethod, "Duration:", time.Since(start), "Error:", err)
		if err = handleAPIKey(ctx, apiKey); err != nil {
			return nil, err
		}
		h, err := handler(ctx, req)
		return h, err
	}
}

func handleAPIKey(ctx context.Context, serverAPIKey string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "metadata_failure")
	}
	apiKeyHeaderValue, ok := md[apiKeyHeaderKey]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "api_key_not_present")
	}
	requestAPIKey := apiKeyHeaderValue[0]
	if requestAPIKey != serverAPIKey {
		return status.Errorf(codes.Unauthenticated, "invalid_api_key")
	}
	return nil
}
