package main

import (
	"context"
	"grpc_driver_hello_go/sidecar"
	"log"
	"net"

	"google.golang.org/grpc"
)

// SidecarServer is the server API for Sidecar service.
// All implementations must embed UnimplementedSidecarServer
// for forward compatibility
type SidecarServer interface {
}

func NewServer() *server {
	return new(server)
}

type server struct {
	sidecar.UnimplementedSidecarServer
}

// 初始化, 主要是为了传配置进去
func (s *server) Init(context.Context, *sidecar.Config) (*sidecar.Response, error) {
	return &sidecar.Response{
		Code: 1,
	}, nil

}

// 启动
func (s *server) Start(context.Context, *sidecar.Request) (*sidecar.Response, error) {
	return &sidecar.Response{
		Code: 1,
	}, nil
}

// 获取状态
func (s *server) Status(context.Context, *sidecar.Request) (*sidecar.Response, error) {
	return &sidecar.Response{
		Code: 1,
	}, nil
}

// 读数据
func (s *server) Read(context.Context, *sidecar.ReadRequest) (*sidecar.ReadResponse, error) {
	return &sidecar.ReadResponse{
		Len:  6,
		Data: []byte{1, 2, 3, 4, 5, 6},
	}, nil
}

func (s *server) Write(context.Context, *sidecar.WriteRequest) (*sidecar.WriteResponse, error) {
	return &sidecar.WriteResponse{
		Code: 1,
	}, nil
}

// 停止
func (s *server) Stop(context.Context, *sidecar.Request) (*sidecar.Response, error) {
	return &sidecar.Response{
		Code: 1,
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("started tcp", "127.0.0.1:8899")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	sidecar.RegisterSidecarServer(grpcServer, NewServer())
	grpcServer.Serve(listener)
}
