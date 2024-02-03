package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/RaniAgus/go-grpc-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func GetRequest(in *pb.CalculationRequestWithOperation) *pb.CalculationRequest {
	return &pb.CalculationRequest{
		A: in.GetA(), B: in.GetB(),
	}
}

func (cs *calculatorServer) Calculate(
	ctx context.Context,
	in *pb.CalculationRequestWithOperation,
) (*pb.CalculationResponse, error) {
	switch in.GetOperation() {
	case pb.Operation_Add:
		return cs.Add(ctx, GetRequest(in))
	case pb.Operation_Substract:
		return cs.Subtract(ctx, GetRequest(in))
	case pb.Operation_Multiply:
		return cs.Multiply(ctx, GetRequest(in))
	case pb.Operation_Divide:
		return cs.Divide(ctx, GetRequest(in))
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown operation")
	}
}

func (cs *calculatorServer) Add(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.GetA() + in.GetB(),
	}, nil
}

func (cs *calculatorServer) Subtract(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.GetA() - in.GetB(),
	}, nil
}

func (cs *calculatorServer) Multiply(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.GetA() * in.GetB(),
	}, nil
}

func (cs *calculatorServer) Divide(
	ctx context.Context,
	in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	if in.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "cannot divide by zero")
	}

	return &pb.CalculationResponse{
		Result: in.GetA() / in.GetB(),
	}, nil
}

func (cs *calculatorServer) Sum(
	ctx context.Context,
	in *pb.NumberStream,
) (*pb.CalculationResponse, error) {
	var sum int64 = 0
	for _, v := range in.GetNumbers() {
		sum += v
	}
	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	server := grpc.NewServer()
	reflection.Register(server) // This is for grpcui

	pb.RegisterCalculatorServer(server, &calculatorServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
