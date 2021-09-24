module ztests/grpcserver

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/grpc v1.40.0
	oldjon.com/protocmd v0.0.0
)

replace oldjon.com/protocmd => ../protocmd
