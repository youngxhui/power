package main

import (
	"github.com/youngxhui/power/core"
	"github.com/youngxhui/power/example/proto"
	"github.com/youngxhui/power/register"
)

func main() {

	power := core.Power{
		Port:       1234,
		ServerName: "Demo1",
	}

	c := core.Config{
		Register: register.Consul{
			Port:    0,
			Address: "127.0.0.1",
		},
	}
	// config := core.DefaultConfig()

	s := power.NewServer(c)
	// 注册proto
	proto.RegisterGreeterServer(s, &GreeterService{})
	// 启动服务
	power.Run()
}
