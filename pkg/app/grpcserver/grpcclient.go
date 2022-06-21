package grpcserver

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/keepchen/app-template/pkg/app/grpcserver/config"
	"github.com/keepchen/app-template/pkg/common/grpc/pb"
	"google.golang.org/grpc"
)

func RunClient(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx,
		fmt.Sprintf("localhost%s", cfg.GrpcServer.Addr),
		grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Printf("connect grpc server error: %v", err)
		return
	}

	time.Sleep(2 * time.Second)
	start := time.Now().Unix()
	fmt.Println("----------start---------", start)
	wg := &sync.WaitGroup{}
	client := pb.NewGreeterClient(conn)
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(loop int, wg *sync.WaitGroup) {
			resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "keepchen"})
			fmt.Println("current:", loop, resp, err)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("----------end---------", end)
	fmt.Println("----------cost---------", end-start)
}
