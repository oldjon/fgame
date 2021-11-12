package main

import (
	"flag"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"oldjon.com/com"
	"oldjon.com/env"
	"oldjon.com/glog"
	"oldjon.com/server"
)

type LoginServer struct {
	server.Server
	serverid      uint64
	gatewayserver []string
	//TODEL
	gatewaylbtype uint32
	gatewaylb     uint64
}

var loginserver *LoginServer

func LoginServer_GetMe() *LoginServer {
	if loginserver == nil {
		loginserver = &LoginServer{}
		loginserver.Derived = loginserver
	}
	return loginserver
}

func (this *LoginServer) Init() bool {

	pprofport := env.Get("loginserver", "pprofport")
	if pprofport != "" {
		go func() {
			http.ListenAndServe(pprofport, nil)
		}()
	}
	redisaddrs := map[string]string{
		"cache": env.Get("global", "redis_cache"),
		"acc":   env.Get("global", "redis_acc"),
		"lock":  env.Get("global", "redis_lock"),
	}
	if !RedisMgr_GetMe().Init(redisaddrs) {
		glog.Error("[启动] 连接redis cluster失败 ")
		return false
	}

	if !this.StartHttpServer() {
		glog.Error("[启动] 启动http服务失败 ")
		return false
	}

	if !this.InitGatewayAddr() {
		glog.Error("[启动] 初始化登录服务失败 ")
		return false
	}

	return true
}

func (this *LoginServer) MainLoop() {
	time.Sleep(time.Second)
	return
}

func (this *LoginServer) Final() bool {
	RedisMgr_GetMe().Close()
	return true
}

func (this *LoginServer) Reload() {
	return
}

func (this *LoginServer) GetServerId() uint64 {
	return this.serverid
}

var (
	logfile = flag.String("logfile", "", "Log file name")
	config  = flag.String("config", "config.json", "config file")
)

func (this *LoginServer) StartHttpServer() bool {
	http.HandleFunc("/time", HandleTime)
	http.HandleFunc("/login", HandleLogin)

	listen, err := net.Listen("tcp", env.Get("loginserver", "httpport"))
	if err != nil {
		glog.Error("[启动] http服务端口监听失败 ", err)
		return false
	}

	go http.Serve(listen, nil)

	return true
}

func (this *LoginServer) InitGatewayAddr() bool {

	gwaddrs := env.Get("loginserver", "gatewayaddrs")
	if gwaddrs == "" {
		return false
	}
	gwlbtype := env.Get("loginserver", "gatewaylbtype")
	if gwlbtype == "" {
		return false
	}

	this.gatewayserver = strings.Split(gwaddrs, "|")
	this.gatewaylbtype = com.StrToUint32(gwlbtype)

	return true
}

func (this *LoginServer) GetGatewayServer(userid uint64) string {
	if len(this.gatewayserver) == 0 {
		return ""
	}
	switch this.gatewaylb {
	case 1:
		index := atomic.AddUint64(&this.gatewaylb, 1)
		return this.gatewayserver[index%uint64(len(this.gatewayserver))]
	default:
		return this.gatewayserver[userid%uint64(len(this.gatewayserver))]
	}
}

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
		glog.SetLogFile(env.Get("loginserver", "log"))
	}

	defer glog.Flush()
	LoginServer_GetMe().Main()
}
