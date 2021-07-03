package initialize

import (
	"fmt"

	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/pkg/tracer"
)

func InitTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		global.ServerConfig.JaegerInfo.Name,
		fmt.Sprintf("%s:%d", global.ServerConfig.JaegerInfo.Host, global.ServerConfig.JaegerInfo.Port))
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
