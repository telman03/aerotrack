package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

type VehicleStatus struct {
	VehicleId 	int64		`json:"vehicle_id"`
	Latitude 	string		`json:"latitude"`
	Longtitude	string		`json:"longtitude"`
	Speed		float64		`json:"speed"`
	Timestamp 	int64		`json:"timestamp"`  
}

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	sub := rdb.Subscribe(ctx, "vehicle_updates")
	defer sub.Close()

	vehicleMap := make(map[int64]VehicleStatus)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🚦 Real-Time Vehicle Dashboard Started...")
	fmt.Println(strings.Repeat("=", 50))


	go func() {
		for msg := range sub.Channel() {
			var status VehicleStatus
			if err := json.Unmarshal([]byte(msg.Payload), &status); err != nil {
				log.Println("❌ JSON decode error:", err)
				continue
			}
			vehicleMap[status.VehicleId] = status
			printTable(vehicleMap)
		}
	}()
	<-sigs
	fmt.Println("\n👋 Shutting down dashboard.")

}
func printTable(data map[int64]VehicleStatus) {
	// Clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%-10s %-12s %-12s %-8s %-20s\n", "VehicleID", "Latitude", "Longitude", "Speed", "Timestamp")
	fmt.Println(strings.Repeat("-", 65))

	for _, v := range data {
		t := time.Unix(v.Timestamp, 0).Format("15:04:05")
		speedColor := getSpeedColor(v.Speed)
		fmt.Printf("%-10d %-12s %-12s %s%-8.2f\033[0m %-20s\n",
			v.VehicleId, v.Latitude, v.Longtitude, speedColor, v.Speed, t)
	}
}

func getSpeedColor(speed float64) string {
	switch {
	case speed < 100:
		return "\033[32m" // green
	case speed < 500:
		return "\033[33m" // yellow
	default:
		return "\033[31m" // red
	}
}
