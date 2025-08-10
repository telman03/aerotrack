package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/telman03/aerotrack/aerotrack/proto"

	"google.golang.org/grpc"
)	

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
        log.Fatalf("Could not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewTrackerServiceClient(conn)

    rand.Seed(time.Now().UnixNano())
    vehicleID := int64(1)

    for {
        update := &pb.LocationUpdate{
            VehicleId: vehicleID,
            Latitude:  randomCoord(40.0, 41.0),
            Longitude: randomCoord(-74.0, -73.0),
            Speed:     rand.Float64()*900 + 100, // 100–1000 km/h
            Timestamp: time.Now().Unix(),
        }

        _, err := client.UpdateLocation(context.Background(), update)
        if err != nil {
            log.Fatalf("Failed to send update: %v", err)
        }

        log.Printf("Sent location update: %+v", update)
        time.Sleep(2 * time.Second)
    }
}

func randomCoord(min, max float64) string {
    return fmt.Sprintf("%.6f", min+(rand.Float64()*(max-min)))
}