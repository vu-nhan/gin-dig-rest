package transport

import (
	"context"
	pb "github.com/vu-nhan/gin-dig-rest/pb/generated"
)

type GrpcServer struct{
	pb.UnimplementedGreeterServer
}

func (s *GrpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}


