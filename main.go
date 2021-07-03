package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/initialize"
	"github.com/childelins/go-grpc-srv/pkg/app"
)

var serviceId string

func init() {
	var err error
	if err = initialize.InitConfig(); err != nil {
		log.Fatalf("初始化配置中心失败: %v", err)
	}
	if err = initialize.InitLogger(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
	serviceId = app.UUID()
	if err = initialize.InitRegistry(serviceId); err != nil {
		log.Fatalf("初始化注册中心失败: %v", err)
	}
	if err = initialize.InitTracer(); err != nil {
		log.Fatalf("初始化链路追踪失败: %v", err)
	}
	if err = initialize.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "", "grpc server ip")
	flag.IntVar(&port, "port", 0, "grpc server port")
	flag.Parse()

	if len(host) > 0 {
		global.ServerConfig.Host = host
	}
	if port > 0 {
		global.ServerConfig.Port = port
	}

	log.Printf("启动gRPC服务[%s], host: %s, port: %d",
		global.ServerConfig.Name,
		global.ServerConfig.Host,
		global.ServerConfig.Port)
	err := initialize.InitGRPCServer(global.ServerConfig.Host, global.ServerConfig.Port)
	if err != nil {
		log.Fatalf("启动gRPC服务失败: %v", err)
	}

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("开始关闭gRPC服务...")

	log.Println("开始注销consul服务...")
	if err := global.Registry.DeRegister(serviceId); err != nil {
		log.Fatalf("注销consul服务[%s]失败: %v", serviceId, err)
	}
	log.Printf("注销consul服务[%s]成功", serviceId)

	log.Println("所有服务已正常退出")
}
