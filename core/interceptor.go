package core

import (
	"context"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/youngxhui/power/log"
	"google.golang.org/grpc"
)

func InterceptorChain() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
		grpcctxtags.UnaryServerInterceptor(),
		grpcmiddleware.ChainUnaryServer(loggingInterceptor),
	))
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info("gRPC 请求方法：", info.FullMethod, "参数:", req)
	resp, err := handler(ctx, req)
	log.Info("gRPC 响应方法:", info.FullMethod, "参数:", resp)
	return resp, err
}
