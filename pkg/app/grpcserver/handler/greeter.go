package handler

import (
	"context"
	"fmt"

	"github.com/keepchen/app-template/pkg/common/grpc/pb"
)

func (*GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", request.Name),
	}, nil
}
