package router

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/childelins/go-grpc-srv/grpc/controller"
	"github.com/childelins/go-grpc-srv/pkg/registry/consul"
	"github.com/childelins/go-grpc-srv/proto"
)

// 注册处理函数
func RegisterServer(g *grpc.Server) {
	// 健康检查
	grpc_health_v1.RegisterHealthServer(g, &consul.HealthServicer{})

	// rpc服务
	proto.RegisterLecturerServer(g, &controller.Lecturer{})
}
