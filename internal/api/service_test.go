package api

import (
	"context"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatherAPI_Add(t *testing.T) {
	service := MatherAPI{}
	testRequest := &MathRequest{
		A: 1,
		B: 3,
	}
	expectedResponse := &MathIntResponse{
		Result: 4,
	}
	result, err := service.Add(context.Background(), testRequest)
	require.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
}

func TestMatherAPI_Subtract(t *testing.T) {
	service := MatherAPI{}
	testRequest := &MathRequest{
		A: 1,
		B: 3,
	}
	expectedResponse := &MathIntResponse{
		Result: -2,
	}
	result, err := service.Subtract(context.Background(), testRequest)
	require.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
}

func TestMatherAPI_Multiply(t *testing.T) {
	service := MatherAPI{}
	testRequest := &MathRequest{
		A: 1,
		B: 3,
	}
	expectedResponse := &MathIntResponse{
		Result: 3,
	}
	result, err := service.Multiply(context.Background(), testRequest)
	require.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
}

func TestMatherAPI_Divide(t *testing.T) {
	service := MatherAPI{}
	t.Run("Happy path", func(t *testing.T) {
		testRequest := &MathRequest{
			A: 1,
			B: 2,
		}
		expectedResponse := &MathFloatResponse{
			Result: 0.5,
		}
		result, err := service.Divide(context.Background(), testRequest)
		require.NoError(t, err)
		assert.Equal(t, expectedResponse, result)
	})
	t.Run("Zero division error", func(t *testing.T) {
		testRequest := &MathRequest{
			A: 1,
			B: 0,
		}
		_, err := service.Divide(context.Background(), testRequest)
		require.Error(t, err)
		assert.Equal(t, ErrZeroDivision, err)
	})
}

func TestMatherAPI_Hello(t *testing.T) {
	service := MatherAPI{}
	t.Run("With name in request", func(t *testing.T) {
		testRequest := &NameRequest{
			Name: "some-name",
		}
		expectedResponse := &HelloResponse{
			Greeting: "Hello some-name",
		}
		result, err := service.Hello(context.Background(), testRequest)
		require.NoError(t, err)
		assert.Equal(t, expectedResponse, result)
	})
	t.Run("Without name in request", func(t *testing.T) {
		testRequest := &NameRequest{}
		expectedResponse := &HelloResponse{
			Greeting: "Hello friend",
		}
		res, err := service.Hello(context.Background(), testRequest)
		require.NoError(t, err)
		assert.Equal(t, expectedResponse, res)
	})
}
