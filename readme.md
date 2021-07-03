# 项目依赖库

- [viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)
- [lumberjack](https://github.com/natefinch/lumberjack)
- [gorm](https://github.com/go-gorm/gorm)
- [opentracing](https://github.com/opentracing/opentracing-go)
- [jeager](https://github.com/uber/jaeger-client-go)
- [consul](https://github.com/hashicorp/consul)
- [grpc-consul-resolver](https://github.com/mbobakov/grpc-consul-resolver)
- [nacos](https://github.com/nacos-group/nacos-sdk-go)
- [uuid](https://github.com/satori/go.uuid)

# 编译 proto 文件
```
// 单个proto文件
protoc --go_out=plugins=grpc:. hello.proto

// 目录下所有proto文件
protoc --go_out=plugins=grpc:. *.proto
```