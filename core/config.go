package core

import "github.com/youngxhui/power/register"

type Config struct {

	// 是否需要服务注册
	IsNeedRegister bool

	// 服务注册中心
	Register register.Register

	// 负载均衡

	// 链路跟踪
}

// DefaultConfig 默认配置
func DefaultConfig() Config {
	return Config{
		IsNeedRegister: true,
		Register: register.Consul{
			Port:    0,
			Address: "127.0.0.1:1234",
		},
	}
}
