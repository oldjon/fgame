package main

import (
	"time"

	"oldjon.com/com"
	"oldjon.com/fcmd"
	"oldjon.com/glog"
	"oldjon.com/server"
)

type ServiceTask struct {
	server.TCPTask
	client  *server.TCPClient
	id      int64
	adddr   string
	close   bool
	service fcmd.Service
	buf     []byte
}

func NewServiceTask(id int64, service fcmd.Service, addr string) *ServiceTask {
	sstask := &ServiceTask{
		TCPTask: *server.NewTCPTask(nil),
		client:  &server.TCPClient{},
		id:      id,
		adddr:   addr,
		buf:     make([]byte, 10240),
	}
	sstask.Derived = sstask
	return sstask
}

func (this *ServiceTask) Send(userid uint64, data []byte) {
	ok := this.SendBytes(com.MarshalUserId(userid, data))
	if !ok {
		glog.Error("[网关] 向服务发送数据失败 ", this.service, ",", this.RemoteAddr(), ",", userid, ",", len(data))
	}
	return
}

func (this *ServiceTask) Connect() bool {
	conn, err := this.client.Connect(this.adddr)
	if err != nil {
		glog.Error("[启动] 连接服务失败 ", this.service, ",", this.adddr)
		return false
	}

	this.Conn = conn
	this.Start()
	this.Verify()
	ServiceMgr_GetMe().AddServiceNode(this)
	glog.Info("[启动] 连接服务成功 ", this.service, ",", this.adddr, ",", this.id)
	return true
}

func (this *ServiceTask) OnClose() {
	this.Reset()
	ServiceMgr_GetMe().RemoveServiceNode(this)
	glog.Info("[网关] 与服务断开连接 ", this.service, ",", this.adddr, ",", this.id, ",", this.close)
	for !this.close {
		glog.Info("[网关] 尝试重连 ", this.service, ",", this.adddr, ",", this.id, ",", this.close)
		if this.Connect() {
			break
		}
		time.Sleep(time.Second * 3)
	}
	return
}

func (this *ServiceTask) ParseMsg(data []byte) bool {
	l := len(data)
	if len(this.buf) < l {
		this.buf = make([]byte, l)
	}
	copy(this.buf, data)

	userid, out, ok := com.UnmarshalUserId(this.buf)
	if !ok {
		glog.Error("[网关] 无玩家id ", this.buf[:l])
		return false
	}

	return UserTaskMgr_GetMe().SendBytes(userid, out)
}
