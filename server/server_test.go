package main

import (
	"context"
	"fmt"
	"sync"
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

func TestConvertCurrencyInvalidSourceCurrency(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "INVALID",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, fmt.Errorf("conversion rate not found for INVALID"))

	res, err := mockServer.Convert(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyInvalidTargetCurrency(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "INVALID",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, fmt.Errorf("conversion rate not found for INVALID"))

	res, err := mockServer.Convert(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyDatabaseError(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}
	mockServer.On("Convert", ctx, req).Return(expectedResp, fmt.Errorf("database error"))

	res, err := mockServer.Convert(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyRateNotFound(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "XYZ",
		TargetCurrency: "INR",
	}

	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}

	mockServer.On("Convert", ctx, req).Return(expectedResp, fmt.Errorf("conversion rate not found for %s", req.SourceCurrency))

	res, err := mockServer.Convert(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, expectedResp, res)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencyTargetNotFound(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "INR",
		TargetCurrency: "UEF",
	}

	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 0,
	}

	mockServer.On("Convert", ctx, req).Return(expectedResp, fmt.Errorf("conversion rate not found for %s", req.SourceCurrency))

	res, err := mockServer.Convert(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, expectedResp, res)
	mockServer.AssertExpectations(t)
}
func TestConvertCurrencyLargeAmount(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ConvertRequest{
		Amount:         1e12,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	expectedResp := &pb.ConvertResponse{
		ConvertedAmount: 74000000000000,
	}

	mockServer.On("Convert", ctx, req).Return(expectedResp, nil)

	res, err := mockServer.Convert(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, res.ConvertedAmount)
	mockServer.AssertExpectations(t)
}

func TestConvertCurrencySimultaneousRequests(t *testing.T) {
	mockServer := new(MockCurrencyConverterServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req1 := &pb.ConvertRequest{
		Amount:         100,
		SourceCurrency: "USD",
		TargetCurrency: "INR",
	}
	req2 := &pb.ConvertRequest{
		Amount:         200,
		SourceCurrency: "EUR",
		TargetCurrency: "INR",
	}
	expectedResp1 := &pb.ConvertResponse{
		ConvertedAmount: 7400,
	}
	expectedResp2 := &pb.ConvertResponse{
		ConvertedAmount: 18302,
	}

	mockServer.On("Convert", ctx, req1).Return(expectedResp1, nil)
	mockServer.On("Convert", ctx, req2).Return(expectedResp2, nil)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		res, err := mockServer.Convert(ctx, req1)
		assert.NoError(t, err)
		assert.Equal(t, expectedResp1.ConvertedAmount, res.ConvertedAmount)
	}()

	go func() {
		defer wg.Done()
		res, err := mockServer.Convert(ctx, req2)
		assert.NoError(t, err)
		assert.Equal(t, expectedResp2.ConvertedAmount, res.ConvertedAmount)
	}()

	wg.Wait()
	mockServer.AssertExpectations(t)
}
