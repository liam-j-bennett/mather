//go:generate protoc --go_out=plugins=grpc:. --proto_path=../../docs/ mather.proto
//go:generate mockgen -source=mather.pb.go -destination=rpc_client_mock.go -package=api -self_package=mather/internal/api MatherClient

package api

import (
	"context"
	"errors"
	"fmt"
)

var ErrZeroDivision = errors.New("zero is an invalid second number as it leads to infinite division")

type MatherAPI struct {
}

func (ma *MatherAPI) Add(ctx context.Context, mr *MathRequest) (*MathIntResponse, error) {
	sum := mr.A + mr.B
	return &MathIntResponse{
		Result: sum,
	}, nil
}

func (ma *MatherAPI) Subtract(ctx context.Context, mr *MathRequest) (*MathIntResponse, error) {
	difference := mr.A - mr.B
	return &MathIntResponse{
		Result: difference,
	}, nil
}

func (ma *MatherAPI) Multiply(ctx context.Context, mr *MathRequest) (*MathIntResponse, error) {
	product := mr.A * mr.B
	return &MathIntResponse{
		Result: product,
	}, nil
}

func (ma *MatherAPI) Divide(ctx context.Context, mr *MathRequest) (*MathFloatResponse, error) {
	if mr.B == 0 {
		return nil, ErrZeroDivision
	}
	product := float64(mr.A) / float64(mr.B)
	return &MathFloatResponse{
		Result: product,
	}, nil
}

func (ma *MatherAPI) Hello(ctx context.Context, mr *NameRequest) (*HelloResponse, error) {
	if mr.Name == "" {
		mr.Name = "friend"
	}
	return &HelloResponse{
		Greeting: greetingCreator(mr.Name),
	}, nil
}

func (ma *MatherAPI) Ready() bool {
	// Usually, we would have dependencies to wait for, but we have nothing
	// ma.someDependency.Ready()
	return true
}

func (ma *MatherAPI) Live() bool {
	// Usually, we would have dependencies we rely on, but we have nothing
	// ma.someDependency.Ready()
	return true
}

func greetingCreator(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
