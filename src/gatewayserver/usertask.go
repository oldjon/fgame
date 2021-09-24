package main

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"oldjon.com/com"
	"oldjon.com/fcmd"
	"oldjon.com/glog"
	"oldjon.com/server"
)

const (
	AUTH_STEP_INIT   = 0
	Task_Max_TimeOut = 60
)

var (
	AesKey = "oldjon.com"
)

type UserTask struct {
	sync.Mutex
	server.TCPTask
	Id       uint64
	Account  string
	AuthKey  string
	AuthStep uint32
	udata    *com.TokenData
	//aesdec     AesDecrypt //TODO
	//aesenc     AesEncrypt //TODO
	activeTime int64
	BeKick     int32
	onlinetime int64
}

func NewUserTask(conn net.Conn) *UserTask {
	nowunix := time.Now().Unix()
	s := &UserTask{
		TCPTask:    *server.NewTCPTask(conn),
		AuthStep:   AUTH_STEP_INIT,
		activeTime: nowunix,
	}
	s.Derived = s
	return s
}

func (this *UserTask) OnClose() {
	if !this.IsVerified() {
		return
	}
	bekick := atomic.LoadInt32(&this.BeKick)
	if bekick == 0 {
		UserTaskMgr_GetMe().Remove(this)
		this.Offline()
	} else if bekick == 2 {
		UserTaskMgr_GetMe().Remove(this)
	}

	glog.Info("[注销] 玩家下线 ", this.Id, ",", this.udata.Platform, ",", time.Now().Unix()-this.onlinetime, ",", this.RemoteAddr())
	return
}

func (this *UserTask) DoVerify(data []byte) bool {
	service := com.GetService(data)
	cmd := com.GetUCmd(data)
	seqid := com.GetSeqId(data)
	if service != fcmd.Service_Gateway || cmd != fcmd.UCmd_GatewayLogin {
		glog.Error("[登录] 只接受验证登录指令 ", this.Conn.RemoteAddr(), ",", service, ",", cmd)
		this.SendErrorMsg(seqid, com.Err_Decode)
		return false
	}
	req := &fcmd.ReqGatewayLogin{}
	_, ok := com.DecodeCmd(data, req)
	if !ok {
		glog.Error("[登录] 反序列化失败 ", this.Conn.RemoteAddr())
		this.SendErrorMsg(seqid, com.Err_Decode)
		return false
	}
	token := &com.TokenData{}
	if !RedisMgr_GetMe().LoadObject(req.Key, token) {
		glog.Error("[登录] 认证失败 ", this.Conn.RemoteAddr(), ",", req.Key)
		return false
	}
	this.Id = token.Id
	this.udata = token
	this.Active()
	this.Verify()

	if oldtask := UserTaskMgr_GetMe().Replace(this); oldtask != nil {
		atomic.StoreInt32(&oldtask.BeKick, 1)
		oldtask.SendErrorMsg(seqid, com.Err_ReLogin)
		oldtask.Stop()
		glog.Info("[登录] 重复登录 ", this.RemoteAddr(), ",", oldtask.Id)
	}
	this.SendBytes(com.EncodeError(seqid, com.Err_OK, nil))
	this.Online()
	glog.Info("[登录] 认证通过 ", this.RemoteAddr(), ",", this.Id)
	return true
}

func (this *UserTask) ParseMsg(data []byte) bool {
	if len(data) < 4 {
		glog.Error("[登录] 消息错误 ", this.Conn.RemoteAddr(), ",", data)
		return false
	}
	if !this.IsVerified() {
		if !this.DoVerify(data) {
			return false
		}
		return true
	}
	this.Active()
	service := com.GetService(data)

	if service == fcmd.Service_Gateway {
		ucmd := com.GetUCmd(data)
		seqId := com.GetSeqId(data)
		if ucmd == fcmd.UCmd_HeartBeat {
			this.SendErrorMsg(seqId, com.Err_OK)
		} else {
			this.SendErrorMsg(seqId, com.Err_Param)
		}
		return true
	}

	//一致性hash到后续节点
	servicenode := ServiceMgr_GetMe().GetService(service, this.Id)
	if servicenode == nil {
		glog.Error("[网关] 不支持的服务类型 ", this.RemoteAddr(), ",", service)
		return false
	}
	servicenode.Send(this.Id, data)
	return true
}

func (this *UserTask) Active() {
	this.activeTime = time.Now().Unix()
	return
}

func (this *UserTask) SendErrorMsg(seqId uint32, errcode uint32) bool {
	return this.SendBytes(com.EncodeError(seqId, errcode, nil))
}

func (this *UserTask) Online() bool {
	this.onlinetime = time.Now().Unix()
	//通知到lobbyserver,设置登录状态等 //TODO
	//通知push上线

	ServiceMgr_GetMe().UserOnline(this.Id)

	return true
}

func (this *UserTask) Offline() {
	//通知lobbyserver,设置下线状态等 //TODO
	//通知push下线

	ServiceMgr_GetMe().UserOffline(this.Id)

	return
}
