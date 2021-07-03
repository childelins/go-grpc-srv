package interceptor

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/childelins/go-grpc-srv/global"
	myTracer "github.com/childelins/go-grpc-srv/pkg/tracer"
)

func OpenTracingServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		// 服务端拦截器则是在MD中把 span提取出来
		spanContext, err := tracer.Extract(opentracing.TextMap, myTracer.MDReaderWriter{MD: md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			global.Logger.Errorf("extract from metadata error: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()
			ctx = opentracing.ContextWithSpan(ctx, span)

			// 设置grom context
			global.DB = global.DB.WithContext(ctx)
		}

		return handler(ctx, req)
	}
}
