package main

import (
	"context"
	"testing"
	"time"

	pb "CurrencyConverter/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCurrencyConverterServer is a mock implementation of the CurrencyConverterServer interface
type MockCurrencyConverterServer struct {
	mock.Mock
}

func (m *MockCurrencyConverterServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

func TestConvertCurrencyToINR(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 7400,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyToUSDFromINR(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         7400,
		SourceCurrency: "INR",
		TargetCurrency: "USD",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 100,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyToEURFromUSD(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "EUR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 100,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyToUSDFromEUR(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "EUR",
		TargetCurrency: "USD",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 100,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencySameCurrency(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "USD",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 100,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyZeroAmount(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         0,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyNegativeAmount(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         -100,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: -7400,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestDefaultCurrencyTypeIsINR(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 7400,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}
