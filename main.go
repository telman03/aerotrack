package main

import (
	"context"
	"log"
	"net"

	pb "github.com/telman03/aerotrack/aerotrack/proto"
	"github.com/telman03/aerotrack/tracker"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)


func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	pb.RegisterTrackerServiceServer(grpcServer, tracker.NewTrackerServer(rdb))
	
	log.Println("TrackerService gRPC server running on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}