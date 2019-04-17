package client

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mather/internal/api"
	"testing"
)

func TestRemoteClient_Add(t *testing.T) {
	arg1 := 1
	arg2 := 2
	expectedRequest := &api.MathRequest{
		A: int64(arg1),
		B: int64(arg2),
	}
	t.Run("Test we get err if client returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		returnError := errors.New("some-error")
		mockClient.EXPECT().Add(gomock.Any(), expectedRequest).Return(nil, returnError)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		_, err := rClient.Add(arg1, arg2)
		assert.Error(t, returnError, err)
	})
	t.Run("Test we get result if client returns ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		expectedResult := &api.MathIntResponse{
			Result: 3,
		}
		mockClient.EXPECT().Add(gomock.Any(), expectedRequest).Return(expectedResult, nil)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		res, err := rClient.Add(arg1, arg2)
		require.NoError(t, err)
		assert.Equal(t, int(expectedResult.Result), res)
	})
}

func TestRemoteClient_Subtract(t *testing.T) {
	arg1 := 1
	arg2 := 2
	expectedRequest := &api.MathRequest{
		A: int64(arg1),
		B: int64(arg2),
	}
	t.Run("Test we get err if client returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		returnError := errors.New("some-error")
		mockClient.EXPECT().Subtract(gomock.Any(), expectedRequest).Return(nil, returnError)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		_, err := rClient.Subtract(arg1, arg2)
		assert.Error(t, returnError, err)
	})
	t.Run("Test we get result if client returns ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		expectedResult := &api.MathIntResponse{
			Result: -1,
		}
		mockClient.EXPECT().Subtract(gomock.Any(), expectedRequest).Return(expectedResult, nil)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		res, err := rClient.Subtract(arg1, arg2)
		require.NoError(t, err)
		assert.Equal(t, int(expectedResult.Result), res)
	})
}

func TestRemoteClient_Multiply(t *testing.T) {
	arg1 := 1
	arg2 := 2
	expectedRequest := &api.MathRequest{
		A: int64(arg1),
		B: int64(arg2),
	}
	t.Run("Test we get err if client returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		returnError := errors.New("some-error")
		mockClient.EXPECT().Multiply(gomock.Any(), expectedRequest).Return(nil, returnError)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		_, err := rClient.Multiply(arg1, arg2)
		assert.Error(t, returnError, err)
	})
	t.Run("Test we get result if client returns ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		expectedResult := &api.MathIntResponse{
			Result: 2,
		}
		mockClient.EXPECT().Multiply(gomock.Any(), expectedRequest).Return(expectedResult, nil)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		res, err := rClient.Multiply(arg1, arg2)
		require.NoError(t, err)
		assert.Equal(t, int(expectedResult.Result), res)
	})
}

func TestRemoteClient_Divide(t *testing.T) {
	arg1 := 1
	arg2 := 2
	expectedRequest := &api.MathRequest{
		A: int64(arg1),
		B: int64(arg2),
	}
	t.Run("Test we get err if client returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		returnError := errors.New("some-error")
		mockClient.EXPECT().Divide(gomock.Any(), expectedRequest).Return(nil, returnError)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		_, err := rClient.Divide(arg1, arg2)
		assert.Error(t, returnError, err)
	})
	t.Run("Test we get result if client returns ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockClient := api.NewMockMatherClient(ctrl)
		expectedResult := &api.MathFloatResponse{
			Result: 0.5,
		}
		mockClient.EXPECT().Divide(gomock.Any(), expectedRequest).Return(expectedResult, nil)
		rClient := RemoteClient{
			RPCClient: mockClient,
		}
		res, err := rClient.Divide(arg1, arg2)
		require.NoError(t, err)
		assert.Equal(t, expectedResult.Result, res)
	})
}
