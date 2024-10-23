package currency

import (
	"context"
	"testing"

	pb "CurrencyConverter/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
	pb.UnimplementedCurrencyConverterServer
}

func (m *MockServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

func TestConvertCurrencyUSDToINR(t *testing.T) {
	result := convertCurrency(100, "USD", "INR")
	assert.Equal(t, 8408.0, result)
}

func TestConvertCurrencyINRToUSD(t *testing.T) {
	result := convertCurrency(7400, "INR", "USD")
	assert.Equal(t, 88.01141769743101, result)
}

func TestConvertCurrencyUSDToEUR(t *testing.T) {
	result := convertCurrency(100, "USD", "EUR")
	assert.Equal(t, 91.88066877936836, result)
}

func TestConvertCurrencyEURToUSD(t *testing.T) {
	result := convertCurrency(100, "EUR", "USD")
	assert.Equal(t, 108.83682207421504, result)
}

func TestConvertCurrencyINRToEUR(t *testing.T) {
	result := convertCurrency(100, "INR", "EUR")
	assert.Equal(t, 100/91.51, result)
}

func TestConvertCurrencyEURToINR(t *testing.T) {
	result := convertCurrency(100, "EUR", "INR")
	assert.Equal(t, 100*91.51, result)
}

func TestConvertCurrencyInvalid(t *testing.T) {
	result := convertCurrency(100, "ABC", "XYZ")
	assert.Equal(t, 0.0, result)
}

func TestServerConvert(t *testing.T) {
	mockServer := new(MockServer)
	req := &pb.ConvertRequest{Amount: 100, SourceCurrency: "USD", TargetCurrency: "INR"}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 8408.0}

	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp.ConvertedAmount, resp.ConvertedAmount)

	mockServer.AssertExpectations(t)
}
