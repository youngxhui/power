package main

import (
	"context"
	"github.com/youngxhui/power/balancer"
	"github.com/youngxhui/power/example/proto"
	"github.com/youngxhui/power/log"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn := balancer.Conn("demo")
	defer conn.Close()
	md := metadata.New(map[string]string{})
	md.Append("token", "hello")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "张三"})
	if err != nil {
		panic(err)
	}
	result, err := c.Add(context.Background(), &proto.AddRequest{A: 2, B: 3})
	if err != nil {
		panic(err)
	}

	log.Info(r.Message)
	log.Info(result.Result)
}
