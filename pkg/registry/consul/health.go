package consul

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServicer struct {
}

// Check 实现健康检查接口，这里直接返回健康状态
func (h *HealthServicer) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthServicer) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}
