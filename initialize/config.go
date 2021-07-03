package initialize

import (
	"encoding/json"

	"github.com/childelins/go-grpc-srv/global"
	"github.com/childelins/go-grpc-srv/pkg/config"
	"github.com/childelins/go-grpc-srv/pkg/nacos"
)

func InitConfig() error {
	conf, err := config.NewConfig("config.yaml")
	if err != nil {
		return err
	}

	err = conf.Unmarshal(&global.NacosConfig)
	if err != nil {
		return err
	}

	// 从nacos中读取配置信息
	client, err := nacos.NewConfigClient(global.NacosConfig.Host, global.NacosConfig.Port, global.NacosConfig.Namespace)
	if err != nil {
		return err
	}

	content, err := client.GetConfig(global.NacosConfig.DataId, global.NacosConfig.Group)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		return err
	}

	return nil
}

/*
func InitConfig() error {
	conf, err := config.NewConfig("config.yaml.bak")
	if err != nil {
		return err
	}
	err = conf.Unmarshal(&global.ServerConfig)
	if err != nil {
		return err
	}

	return nil
}
*/
