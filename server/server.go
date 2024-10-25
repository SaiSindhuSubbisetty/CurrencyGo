package main

import (
	pb "CurrencyConverter/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

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

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the CurrencyConverter service with the gRPC server
	pb.RegisterCurrencyConverterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
