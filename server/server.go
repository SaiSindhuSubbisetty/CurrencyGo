package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "CurrencyConverter/proto"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
	db *sql.DB
}

// Initializes a connection to PostgreSQL
func initDB() (*sql.DB, error) {
	connStr := "user=postgres password=1234 dbname=currencydb sslmode=disable" // Updated username
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// convertCurrency retrieves conversion rates from the database
func (s *server) convertCurrency(ctx context.Context, amount float64, sourceCurrency, targetCurrency string) (float64, error) {
	var sourceRate, targetRate float64

	// Retrieve source rate
	err := s.db.QueryRowContext(ctx, "SELECT rate FROM conversion_rates WHERE currency = $1", sourceCurrency).Scan(&sourceRate)
	if err != nil {
		log.Printf("Error retrieving source rate for %s: %v", sourceCurrency, err)
		return 0, fmt.Errorf("conversion rate not found for %s", sourceCurrency)
	}

	// Retrieve target rate
	err = s.db.QueryRowContext(ctx, "SELECT rate FROM conversion_rates WHERE currency = $1", targetCurrency).Scan(&targetRate)
	if err != nil {
		log.Printf("Error retrieving target rate for %s: %v", targetCurrency, err)
		return 0, fmt.Errorf("conversion rate not found for %s", targetCurrency)
	}

	// Convert the amount
	inrAmount := amount * sourceRate
	return inrAmount / targetRate, nil
}

// Convert implements the gRPC method for currency conversion
func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	amount := req.GetAmount()
	sourceCurrency := req.GetSourceCurrency()
	targetCurrency := req.GetTargetCurrency()

	// Call the conversion function
	convertedAmount, err := s.convertCurrency(ctx, amount, sourceCurrency, targetCurrency)
	if err != nil {
		return nil, err
	}

	// Return the response with the converted amount
	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func main() {
	// Initialize the database
	db, err := initDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{db: db})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
