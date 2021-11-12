module oldjon.com/fgame

go 1.17

require (
	oldjon.com/protobuf v0.0.0 // indirect
)

replace (
	oldjon.com/com => ./src/common
	oldjon.com/dbobj => ./src/dbobj
	oldjon.com/env => ./src/base/env
	oldjon.com/fcmd => ./src/fcmd
	oldjon.com/gatewayserver => ./src/gatewayserver
	oldjon.com/glog => ./src/base/glog
	oldjon.com/loginserver => ./src/loginserver
	oldjon.com/protobuf => ./src/base/protobuf
	oldjon.com/redis-go-cluster => ./src/base/redis-go-cluster
	oldjon.com/server => ./src/base/server
)
