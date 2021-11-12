package main

import (
	"flag"
	"net/http"

	_ "net/http/pprof"

	"oldjon.com/env"
	"oldjon.com/glog"
	"oldjon.com/server"
)

type GatewayServer struct {
	server.Server
	serverid  uint64
	tcpserver *server.TCPServer
}

var gatewayserver *GatewayServer

func GatewayServer_GetMe() *GatewayServer {
	if gatewayserver == nil {
		gatewayserver = &GatewayServer{
			tcpserver: &server.TCPServer{},
		}
		gatewayserver.Derived = gatewayserver
	}
	return gatewayserver
}

func (this *GatewayServer) Init() bool { //TODO
	pprofport := env.Get("gatewayserver", "pprofport")
	if pprofport != "" {
		go func() {
			http.ListenAndServe(pprofport, nil)
		}()
	}

	redisaddrs := map[string]string{
		"cache": env.Get("global", "redis_cache"),
		"lock":  env.Get("global", "redis_lock"),
	}

	if !RedisMgr_GetMe().Init(redisaddrs) {
		glog.Error("[启动] 连接redis cluster失败 ")
		return false
	}

	if !ServiceMgr_GetMe().LoadServices() {
		glog.Error("[启动] 连接服务失败 ")
		return false
	}

	if err := this.tcpserver.Bind(env.Get("gatewayserver", "tcpport")); err != nil {
		glog.Error("[启动] 网关服务监听端口失败 ", err)
		return false
	}

	return true
}

func (this *GatewayServer) MainLoop() {

	conn, err := this.tcpserver.Accept()
	if err != nil {
		return
	}
	NewUserTask(conn).Start()
	return
}

func (this *GatewayServer) Reload() {
	return
}

func (this *GatewayServer) GetServerId() uint64 {
	return this.serverid
}

func (this *GatewayServer) Final() bool {
	UserTaskMgr_GetMe().CloseAll()
	this.tcpserver.CLose()
	RedisMgr_GetMe().Close()
	return true
}

var (
	logfile = flag.String("logfile", "", "Log file name")
	config  = flag.String("config", "config.json", "config file")
)

func main() {
	flag.Parse()
	env.Load(*config)
	loglevel := env.Get("global", "loglevel")
	if loglevel != "" {
		flag.Lookup("stderrthreshold").Value.Set(loglevel)
	}

	logtostderr := env.Get("global", "logtostderr")
	if logtostderr != "" {
		flag.Lookup("logtostderr").Value.Set(logtostderr)
	}

	if *logfile != "" {
		glog.SetLogFile(*logfile)
	} else {
		glog.SetLogFile(env.Get("gatewayserver", "log"))
	}

	defer glog.Flush()
	GatewayServer_GetMe().Main()
}
