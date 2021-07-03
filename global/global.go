package global

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"

	"github.com/childelins/go-grpc-srv/config"
	"github.com/childelins/go-grpc-srv/pkg/logger"
	"github.com/childelins/go-grpc-srv/pkg/registry"
)

var (
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig
	Logger       *logger.Logger
	DB           *gorm.DB
	Tracer       opentracing.Tracer
	Registry     registry.Registry
)
