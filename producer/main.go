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
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTrackerServiceClient(conn)


	vehicleIDs := []int64{1, 2, 3, 4, 5}


	for _, id := range vehicleIDs {
		go simulateVehicle(client, id)
	}


	select {}
}

func simulateVehicle(client pb.TrackerServiceClient, vehicleId int64) {
	lat := 40.0 + rand.Float64()*0.5
	lon := 49.0 + rand.Float64()*0.5
	speed := 50.0 + rand.Float64()*20.0


	for {
		lat += (rand.Float64() - 0.5) * 0.01
		lon += (rand.Float64() - 0.5) * 0.01

		speed += (rand.Float64() - 0.5) * 2.0
		if speed < 0 {
			speed = 0
		}

		req := &pb.LocationUpdate{
			VehicleId: vehicleId,
			Latitude: fmt.Sprintf("%.6f", lat),
			Longitude: fmt.Sprintf("%.6f", lon),
			Speed: float64(speed),
			Timestamp: time.Now().Unix(),
		}

		_, err := client.UpdateLocation(context.Background(), req)
		if err != nil {
			log.Printf("[Vehicle %d] Error sending update: %v", vehicleId, err)
		} else {
			log.Printf("[Vehicle %d] Sent update: Lat=%s, Lon=%s, Speed=%.2f", vehicleId, req.Latitude, req.Longitude, req.Speed)
		}

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}

 
}
