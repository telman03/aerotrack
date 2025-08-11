package main

import (
    "context"
    "log"
    "math/rand"
    "time"
	"fmt"
    pb "github.com/telman03/aerotrack/aerotrack/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewTrackerServiceClient(conn)

    rand.Seed(time.Now().UnixNano())

    vehicleIDs := []int64{101, 102, 103, 104} // multiple vehicles

    for {
        for _, id := range vehicleIDs {
            update := &pb.LocationUpdate{
                VehicleId: id,
                Latitude:   randomCoord(40.0, 41.0), // random within Baku-ish area
                Longitude:  randomCoord(49.0, 50.0),
                Speed:      randomSpeed(200, 900),
                Timestamp:  time.Now().Unix(),
            }

            _, err := client.UpdateLocation(context.Background(), update)
            if err != nil {
                log.Printf("❌ Failed to send update for vehicle %d: %v", id, err)
            } else {
                log.Printf("✅ Sent update for vehicle %d", id)
            }
        }

        time.Sleep(1 * time.Second)
    }
}

func randomCoord(min, max float64) string {
    return formatFloat(min + rand.Float64()*(max-min))
}

func randomSpeed(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func formatFloat(f float64) string {
    return fmt.Sprintf("%.6f", f)
}
