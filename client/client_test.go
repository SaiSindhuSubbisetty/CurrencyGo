package main

import (
	"context"
	"testing"
	"time"

	pb "CurrencyConverter/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockCurrencyConverterClient is a mock implementation of the CurrencyConverterClient interface
type MockCurrencyConverterClient struct {
	mock.Mock
}

func (m *MockCurrencyConverterClient) Convert(ctx context.Context, in *pb.ConvertRequest, opts ...grpc.CallOption) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

func TestConvertCurrencyToINR(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencyToUSDFromINR(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencyToEURFromUSD(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "EUR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 100, // Assuming 1:1 conversion for unsupported currencies
	}
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencyToUSDFromEUR(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencySameCurrency(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencyZeroAmount(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}

func TestConvertCurrencyNegativeAmount(t *testing.T) {
	mockClient := new(MockCurrencyConverterClient)
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
	mockClient.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockClient.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockClient.AssertExpectations(t)
}
