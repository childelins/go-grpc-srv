package initialize

import (
	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/pkg/registry/consul"
)

func InitRegistry(serviceId string) error {
	registry := consul.NewRegistry(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	err := registry.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		return err
	}

	global.Registry = registry
	return nil
}
