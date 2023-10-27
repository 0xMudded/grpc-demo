package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	number_generator "github.com/0xMudded/grpc-demo/number-generator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type numberGeneratorServer struct {
	number_generator.UnimplementedNumberGeneratorServer
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func (ngs numberGeneratorServer) Generate(ctx context.Context, req *number_generator.GenerateRequest) (*number_generator.GenerateResponse, error) {
	return &number_generator.GenerateResponse{
		RandomNumber: int64(generateRandomNumber(int(req.StartRange), int(req.EndRange))),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("error creating listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	reflection.Register(serverRegistrar)
	service := &numberGeneratorServer{}
	number_generator.RegisterNumberGeneratorServer(serverRegistrar, service)

	log.Fatal(serverRegistrar.Serve(lis))
}
