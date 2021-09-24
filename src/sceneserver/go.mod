module sceneserver

go 1.14

require (
	oldjon.com/protobuf v0.0.0
	oldjon.com/env v0.0.0
	oldjon.com/glog v0.0.0
	oldjon.com/com v0.0.0
	oldjon.com/fcmd v0.0.0
	github.com/oldjon/gonet v0.0.0-20161121104625-7c2476775e64
)

replace (
	oldjon.com/env => ../base/env
	oldjon.com/glog => ../base/glog
	oldjon.com/com => ../common
	oldjon.com/fcmd => ../fcmd
)
