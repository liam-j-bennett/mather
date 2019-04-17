package client

import (
	"context"
	"google.golang.org/grpc/metadata"
	"mather/internal/api"
	"time"
)

type RemoteClient struct {
	RPCClient api.MatherClient
	APIKey    string
}

// TODO: Move to a Client Interceptor
func setContext(apiKey string) context.Context {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	md := metadata.Pairs("api_key", apiKey)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

func (rc *RemoteClient) Add(a, b int) (int, error) {
	ctx := setContext(rc.APIKey)
	req := &api.MathRequest{
		A: int64(a),
		B: int64(b),
	}
	res, err := rc.RPCClient.Add(ctx, req)
	if err != nil {
		return 0, err
	}
	return int(res.Result), nil
}

func (rc *RemoteClient) Subtract(a, b int) (int, error) {
	ctx := setContext(rc.APIKey)
	req := &api.MathRequest{
		A: int64(a),
		B: int64(b),
	}
	res, err := rc.RPCClient.Subtract(ctx, req)
	if err != nil {
		return 0, err
	}
	return int(res.Result), nil
}

func (rc *RemoteClient) Multiply(a, b int) (int, error) {
	ctx := setContext(rc.APIKey)
	req := &api.MathRequest{
		A: int64(a),
		B: int64(b),
	}
	res, err := rc.RPCClient.Multiply(ctx, req)
	if err != nil {
		return 0, err
	}
	return int(res.Result), nil
}

func (rc *RemoteClient) Divide(a, b int) (float64, error) {
	ctx := setContext(rc.APIKey)
	req := &api.MathRequest{
		A: int64(a),
		B: int64(b),
	}
	res, err := rc.RPCClient.Divide(ctx, req)
	if err != nil {
		return 0, err
	}
	return res.Result, nil
}
