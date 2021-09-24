module gatewayserver

go 1.16

require (
	github.com/golang/protobuf v1.5.0
	oldjon.com/com v0.0.0
	oldjon.com/env v0.0.0
	oldjon.com/fcmd v0.0.0
	oldjon.com/glog v0.0.0
	oldjon.com/redis-go-cluster v0.0.0
	oldjon.com/server v0.0.0
)

replace (
	oldjon.com/com => ../common
	oldjon.com/env => ../base/env
	oldjon.com/fcmd => ../fcmd
	oldjon.com/glog => ../base/glog
	oldjon.com/protobuf => ../base/protobuf
	oldjon.com/redis-go-cluster => ../base/redis-go-cluster
	oldjon.com/server => ../base/server
)
