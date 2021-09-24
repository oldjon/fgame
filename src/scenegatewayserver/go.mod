module server

go 1.16

require (
	github.com/cpp2go/gonet v0.0.0-20161121104625-7c2476775e64 // indirect
	//github.com/oldjon/gonet v0.0.0-20161121104625-7c2476775e64
	oldjon.com/base/env v0.0.0 // indirect
	oldjon.com/base/glog v0.0.0 // indirect
	oldjon.com/common v0.0.0 // indirect
	oldjon.com/fcmd v0.0.0 // indirect
)

replace (
	oldjon.com/base/env => ../base/env
	oldjon.com/base/glog => ../base/glog
	oldjon.com/common => ../common
	oldjon.com/fcmd => ../fcmd
)
