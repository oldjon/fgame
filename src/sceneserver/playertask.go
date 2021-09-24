package main

/*
import (
	"fmt"
	"net"

	"github.com/oldjon/gonet"
	"oldjon.com/common"
)

type PlayerTask struct {
	gonet.TcpTask
}

func NewPlayerTask(conn net.Conn) *PlayerTask {
	s := &PlayerTask{
		TcpTask: *gonet.NewTcpTask(conn),
	}
	s.Derived = s
	return s
}

func (this *PlayerTask) ParseMsg(data []byte, flag byte) bool {

	cmd := usercmd.UserCmd(common.GetCmd(data))

	switch cmd {
	case usercmd.UserCmd_Login:
		{
			revCmd, ok := common.DecodeCmd(data, flag, &usercmd.ReqUserLogin{}).(*usercmd.ReqUserLogin)
			if !ok {
				return false
			}

			fmt.Println("> ", cmd, ",", *revCmd.Account, ",", *revCmd.Password, ",", *revCmd.Key)

			this.Verify()

			retCmd := &usercmd.ReqUserLogin{
				Account:  revCmd.Account,
				Password: revCmd.Password,
				Key:      revCmd.Key,
			}
			this.SendCmd(usercmd.UserCmd_Login, retCmd)
		}

	}

	return true
}

func (this *PlayerTask) SendCmd(cmd usercmd.UserCmd, msg common.Message) bool {
	data, flag, err := common.EncodeCmd(uint16(cmd), msg)
	if err != nil {
		fmt.Println("[服务] 发送失败 cmd:", cmd, ",len:", len(data), ",err:", err)
		return false
	}
	return this.AsyncSend(data, flag)
}

func (this *PlayerTask) OnClose() {

}
*/
