package core

import (
	"github.com/youngxhui/power/log"
	"github.com/youngxhui/power/register"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type Power struct {
	ServerName string
	Port       int
}

var g *grpc.Server

// 服务注册中心
var r register.Register

// NewServer 创建新的服务
func (power Power) NewServer(config Config) *grpc.Server {
	if config.IsNeedRegister  {
		r = config.Register

		r.RegisterServer(power.ServerName, power.Port, "127.0.0.1")
	}

	g = grpc.NewServer()

	return g
}

// Run 启动服务
func (power *Power) Run() {
	l, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(power.Port))

	if err != nil {
		log.Error("服务启动失败")
	}
	log.Info("服务启动", "监听地址 127.0.0.1:"+strconv.Itoa(power.Port))
	if err = g.Serve(l); err != nil {
		log.Error("发生了错误", err.Error())
		panic(err)
	}

	defer r.UnRegisterServer(power.ServerName)
	//power.Register.UnRegisterServer(power.ServerName)
}
