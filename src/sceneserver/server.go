package main

import (
	"flag"
	"fmt"

	"github.com/oldjon/gonet"
	"oldjon.com/base/env"
	"oldjon.com/base/glog"
)

type SceneServer struct {
	gonet.Service
	tcpser *gonet.TcpServer
}

var sceneserver *SceneServer

func GetSceneServer() *SceneServer {
	if sceneserver == nil {
		sceneserver = &SceneServer{
			tcpser: &gonet.TcpServer{},
		}
		sceneserver.Derived = sceneserver
	}
	return sceneserver
}

func (this *SceneServer) Init() bool {

	err := this.tcpser.Bind(":" + env.Get("sceneserver", "tcpport"))
	if err != nil {
		fmt.Println("绑定端口失败")
		return false
	}
	return true
}

func (this *SceneServer) Reload() {
	//重新加载配置逻辑
}

func (this *SceneServer) MainLoop() {
	conn, err := this.tcpser.Accept()
	if err != nil {
		return
	}
	NewGWTask(conn).Start()
}

func (this *SceneServer) Final() bool {
	this.tcpser.Close()
	return true
}

var (
	logfile = flag.String("logfile", "", "log file name")
	config  = flag.String("config", "config.json", "config path")
)

func main() {
	flag.Parse()
	if !env.Load(*config) {
		return
	}
	loglevel := env.Get("global", "loglevel")
	if loglevel != "" {
		flag.Lookup("stderrthreshold").Value.Set(loglevel)
	}
	logtostderr := env.Get("global", "logtostrerr")
	if logtostderr != "" {
		flag.Lookup("logtostderr").Value.Set(loglevel)
	}
	if *logfile != "" {
		glog.SetLogFile(*logfile)
	} else {
		glog.SetLogFile(env.Get("sceneserver", "log"))
	}
	defer glog.Flush()

	GetSceneServer().Main()

	glog.Info("[关闭] 场景服务器关闭")
}
