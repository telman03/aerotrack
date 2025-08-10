package tracker

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/telman03/aerotrack/aerotrack/proto"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrackerServer struct {
	pb.UnimplementedTrackerServiceServer
	rdb *redis.Client
}

func NewTrackerServer(rbd *redis.Client) *TrackerServer {
	return &TrackerServer{rdb: rbd}
}

func (s *TrackerServer) UpdateLocation(ctx context.Context, req *pb.LocationUpdate) (*emptypb.Empty, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("vehicle:%d", req.VehicleId)
	err = s.rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
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