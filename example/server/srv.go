package main

import (
	"Demo/example/proto"
	"context"
)

type GreeterService struct {
}

func (s *GreeterService) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello," + request.Name + " from go"}, nil
}

func (s *GreeterService) Add(ctx context.Context, request *proto.AddRequest) (*proto.AddReply, error) {
	return &proto.AddReply{Result: request.A + request.B}, nil
}