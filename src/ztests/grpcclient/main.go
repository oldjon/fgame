package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	pb "oldjon.com/protocmd"
)

func main() {
	// 连接
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.ReqHello{Value: "gRPC"}
	res, err := c.Hello(context.Background(), req)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Value)
}
