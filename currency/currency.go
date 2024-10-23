package currency

import (
	pb "CurrencyConverter/proto"
	"context"
	"fmt"
)

// Conversion rates relative to INR (base currency)
var conversionRates = map[string]float64{
	"INR": 1.0,
	"USD": 84.08,
	"EUR": 91.51,
}

// Convert converts the amount from the source currency to the target currency
func convertCurrency(amount float64, sourceCurrency, targetCurrency string) float64 {
	// If currency is not provided, default to INR
	if sourceCurrency == "" {
		sourceCurrency = "INR"
	}
	if targetCurrency == "" {
		targetCurrency = "INR"
	}

	// Retrieve the conversion rates
	sourceRate, ok1 := conversionRates[sourceCurrency]
	targetRate, ok2 := conversionRates[targetCurrency]

	// Check if both source and target currencies are valid
	if !ok1 || !ok2 {
		fmt.Printf("Conversion rate not found for %s or %s\n", sourceCurrency, targetCurrency)
		return 0
	}

	// Convert to INR first, then convert to target currency
	inrAmount := amount * sourceRate
	return inrAmount / targetRate
}

// Server implementing the CurrencyConverter gRPC service
type server struct {
	pb.UnimplementedCurrencyConverterServer
}

// Convert implements the gRPC method
func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	amount := req.GetAmount()
	sourceCurrency := req.GetSourceCurrency()
	targetCurrency := req.GetTargetCurrency()

	// Call the conversion function
	convertedAmount := convertCurrency(amount, sourceCurrency, targetCurrency)

	// Return the response with the converted amount
	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}
