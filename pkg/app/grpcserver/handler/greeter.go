package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/keepchen/app-template/pkg/common/grpc/pb"
)

//SayHello hello处理方法
func (*GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	time.Sleep(2 * time.Second)
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", request.Name),
	}, nil
}
