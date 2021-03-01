package balancer

import (
	"google.golang.org/grpc"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

func Conn(SrvName string) *grpc.ClientConn {
	target := "consul://127.0.0.1:8500/"+SrvName

	conn, err := grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		panic(err)
	}

	return conn
}
