package main

import (
	"grpc-server/config"
	student "grpc-server/proto/students"
	"grpc-server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Inisialisasi listener pada port yang ditentukan di config
	lis, err := net.Listen("tcp", "localhost:9991")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Sign up service to gRPC server
	studentService := &service.StudentService{}
	student.RegisterStudentServiceServer(grpcServer, studentService)

	// Running gRPC server
	log.Printf("gRPC server is running on port %s", cfg.AppPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
