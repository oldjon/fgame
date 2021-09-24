package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "oldjon.com/protocmd"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.ReqHello) (*pb.RetHello, error) {
	reply := &pb.RetHello{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
