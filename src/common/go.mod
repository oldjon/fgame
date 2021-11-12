module com

go 1.14

require (
    oldjon.com/protobuf v0.0.0 // indirect
    oldjon.com/glog v0.0.0 // indirect
    oldjon.com/fcmd v0.0.0 // indirect
    oldjon.com/redis-go-cluster v0.0.0
)

replace (
	oldjon.com/protobuf => ../base/protobuf
	oldjon.com/glog => ../base/glog
	oldjon.com/fcmd => ../fcmd
	oldjon.com/redis-go-cluster => ../base/redis-go-cluster
)