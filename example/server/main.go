package main

import (
	"Demo/core"
	"Demo/example/proto"
	"Demo/register"
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
