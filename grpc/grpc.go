package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGet(port int64, server func(c *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//初始化grpc
	s := grpc.NewServer()
	server(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
