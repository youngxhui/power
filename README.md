# POWER  [![Go](https://github.com/youngxhui/power/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/youngxhui/power/actions/workflows/go.yml)[![CodeQL](https://github.com/youngxhui/power/actions/workflows/codeql-analysis.yml/badge.svg?branch=master)](https://github.com/youngxhui/power/actions/workflows/codeql-analysis.yml)

**⚠️WIP⚠️**

开箱即用的 gRPC 框架。

## 服务端

构建简单的服务端

```go
package main

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
```

## 客户端

```go

```

