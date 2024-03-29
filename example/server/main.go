package main

import (
	"github.com/youngxhui/power/core"
	"github.com/youngxhui/power/example/proto"
)

func main() {

	power := core.Power{
		Port:       1234,
		ServerName: "Demo1",
	}

	config := core.DefaultConfig()

	s := power.NewServer(config)
	// 注册proto
	proto.RegisterGreeterServer(s, &GreeterService{})
	// 启动服务
	power.Run()
}
