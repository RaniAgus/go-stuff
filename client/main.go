package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/RaniAgus/go-grpc-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = flag.String("addr", ":50051", "The server address")
)

func main() {
	flag.Parse()

	var creds credentials.TransportCredentials
	if strings.HasPrefix(*address, ":") {
		creds = insecure.NewCredentials()
	} else {
		creds = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *address, opts...)
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *address, err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	nums := []int64{1, 2, 3, 4, 5}

	res, err := client.Sum(ctx, &pb.NumberStream{
		Numbers: nums,
	})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Printf("Sum result: %d", res.GetResult())
}
