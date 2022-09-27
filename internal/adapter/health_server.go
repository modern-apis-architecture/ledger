package adapter

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
)

type HealthServer struct {
}

func (h *HealthServer) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Printf("✅ Server's status is %s", grpc_health_v1.HealthCheckResponse_SERVING)
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthServer) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return nil
}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}
