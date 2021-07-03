package initialize

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/grpc/interceptor"
	"github.com/childelins/go-grpc-srv/router"
)

func InitGRPCServer(address string, port int) error {
	g := grpc.NewServer(grpc.UnaryInterceptor(interceptor.OpenTracingServerInterceptor(global.Tracer)))
	router.RegisterServer(g)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}

	go func() {
		err = g.Serve(listener)
		if err != nil {
			log.Println("gRPC服务异常终止:", err)
		}
	}()

	return nil
}
