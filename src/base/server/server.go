package server

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"

	"oldjon.com/glog"
)

type IServer interface {
	Init() bool
	MainLoop()
	Reload()
	Final() bool
}

type Server struct {
	closed  bool
	Derived IServer
}

func (this *Server) Close() {
	this.closed = true
	return
}

func (this *Server) IsClosed() bool {
	return this.closed
}

func (this *Server) SetCPUNum(num int) {
	if num > 0 {
		runtime.GOMAXPROCS(num)
	} else if num == -1 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	return
}

func (this *Server) Main() bool {
	defer func() {
		this.Derived.Final()
		if err := recover(); err != nil {
			glog.Error("[异常] ", err, "\n", string(debug.Stack()))
		}
		glog.Info("关闭服务器完成")
		glog.Flush()
	}()

	rand.Seed(time.Now().Unix() ^ int64(os.Getpid()))

	if this.Derived == nil {
		glog.Error("[启动] Server Derived为空 ")
		return false
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE, syscall.SIGHUP)
	go func() {
		for sig := range ch {
			switch sig {
			case syscall.SIGHUP:
				glog.Info("[服务] 收到重新加载配置信号！")
				this.Derived.Reload()
			case syscall.SIGPIPE:
			default:
				this.Close()
			}
			glog.Info("[服务] 收到信号 ", sig)
		}
	}()

	this.SetCPUNum(runtime.NumCPU())

	glog.Info("[启动] 开始初始化")

	if !this.Derived.Init() {
		glog.Error("[启动] 初始化失败")
		return false
	}

	glog.Info("[启动] 初始化成功")

	for !this.IsClosed() {
		this.Derived.MainLoop()
	}

	return true
}
