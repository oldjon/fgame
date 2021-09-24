module ztests/grpcclient

go 1.16

require (
	google.golang.org/grpc v1.40.0
	oldjon.com/protocmd v0.0.0
)

replace oldjon.com/protocmd => ../protocmd
