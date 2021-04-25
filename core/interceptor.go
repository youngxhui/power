package core

import (
	"github.com/youngxhui/power/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"context"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/youngxhui/power/log"
	"github.com/youngxhui/power/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var c Config

func InterceptorChain(config Config) grpc.ServerOption {
	c = config
	return grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
		grpcctxtags.UnaryServerInterceptor(),
		grpcmiddleware.ChainUnaryServer(loggingInterceptor),
		grpcmiddleware.ChainUnaryServer(jwtInterceptor),
	))
}

// 日志拦截器
func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info("gRPC 请求方法：", info.FullMethod, "参数:", req)
	resp, err := handler(ctx, req)
	log.Info("gRPC 响应方法:", info.FullMethod, "参数:", resp)
	return resp, err
}

// jwt拦截器
func jwtInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	path := util.GetConfPath()
	yamlFile, err := ioutil.ReadFile(path)
	y := new(config.Yaml)
	err = yaml.Unmarshal(yamlFile, y)
	log.Info(path, yamlFile, y.Security)
	log.Info("jwt ")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "未授权")
	}
	if token, ok := md["jwt"]; ok {
		t := strings.Join(token, "")
		parseToken, err := util.ParseToken(t, c.Jwt.Secret)
		if err != nil {
			log.Waring("没有权限")
		}
		log.Info(parseToken)
	} else {
		// 遍历路径

		return nil, status.Error(codes.Unauthenticated, "未授权1")
	}

	resp, err := handler(ctx, req)
	return resp, err
}
