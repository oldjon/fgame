package main

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/oldjon/gonet"
	"oldjon.com/base/glog"
	"oldjon.com/com"
	"oldjon.com/fcmd"
)

type GWTask struct { //网关
	GWId uint32
	gonet.TcpTask
}

func NewGWTask(conn net.Conn) *GWTask {
	s := &GWTask{
		TcpTask: *gonet.NewTcpTask(conn),
	}
	s.Derived = s
	return s
}

func (this *GWTask) ParseMsg(data []byte, flag byte) bool {

	cmd := fcmd.SCmd(com.GetCmd(data))

	switch cmd {
	case fcmd.SCmd_ServerRegist:
		{
			glog.Info("[网关] 接收到网关链接")
			reqcmd, ok := com.DecodeCmd(data, flag, &fcmd.ReqServerRegist{}).(*fcmd.ReqServerRegist)
			if !ok {
				return false
			}
			if reqcmd.GWId == 0 {
				this.Close()
				glog.Info("[网关] 网关注册失败,GWId错误")
				return false
			}

			this.GWId = reqcmd.GWId
			GWMgr_GetMe().AddGWTask(this)

			retCmd := &fcmd.RetServerRegist{
				State: com.State_Ok,
			}

			this.SendCmd(fcmd.SCmd_ServerRegist, retCmd)
			glog.Info("[网关] 网关注册成功")
		}
	case fcmd.SCmd_Player: //处理玩家指令
		PlayerMgr_GetMe().ParseMsg(data[com.CmdHeaderSize:], flag)
	}

	return true
}

func (this *GWTask) SendCmd(cmd fcmd.SCmd, msg proto.Message) bool {
	data, flag, err := com.EncodeCmd(uint16(cmd), msg)
	if err != nil {
		fmt.Println("[服务] 发送失败 cmd:", cmd, ",len:", len(data), ",err:", err)
		return false
	}
	return this.AsyncSend(data, flag)
}

func (this *GWTask) OnClose() { //关闭链接
}
