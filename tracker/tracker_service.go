package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	pb "github.com/telman03/aerotrack/aerotrack/proto"
)

type TrackerServer struct {
	pb.UnimplementedTrackerServiceServer
	rdb *redis.Client
}

func NewTrackerServer(rbd *redis.Client) *TrackerServer {
	return &TrackerServer{rdb: rbd}
}

func (s *TrackerServer) UpdateLocation(ctx context.Context, req *pb.LocationUpdate) (*pb.UpdateResponse, error) {
	key := fmt.Sprintf("vehicle:%d", req.VehicleId)

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = s.rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return nil, err
	}
	if err := s.rdb.Publish(ctx, "vehicle_updates", data).Err(); err != nil {
		return nil, err
	}

	log.Printf("Updated location for vehicle %d and published to channel", req.VehicleId)

	return &pb.UpdateResponse{Message: "Location updated successfully"}, nil
}

func (s *TrackerServer) GetVehicleStatus(ctx context.Context, req *pb.VehicleID) (*pb.VehicleStatus, error){
	key := fmt.Sprintf("vehicle:%d", req.VehicleId)
	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var status pb.VehicleStatus
	if err := json.Unmarshal([]byte(val), &status); err != nil {
		return nil, err
	}

	return &status, nil
}