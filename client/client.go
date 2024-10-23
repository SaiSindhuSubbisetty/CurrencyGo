package main

import (
	pb "CurrencyConverter/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCurrencyConverterClient(conn)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send a currency conversion request
	r, err := c.Convert(ctx, &pb.ConvertRequest{Amount: 100, SourceCurrency: "USD", TargetCurrency: "INR"})
	if err != nil {
		log.Fatalf("could not convert: %v", err)
	}

	// Log the converted amount
	log.Printf("Converted Amount: %f", r.GetConvertedAmount())
}
