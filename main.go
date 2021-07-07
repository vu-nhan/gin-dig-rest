package main

import (
	pb "github.com/vu-nhan/gin-dig-rest/pb/generated"
	"github.com/vu-nhan/gin-dig-rest/provider"
	"github.com/vu-nhan/gin-dig-rest/transport"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main()  {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}


	iocProvider := provider.IocProvider{}
	iocProvider.InitIocProvider()

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &transport.GrpcServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}