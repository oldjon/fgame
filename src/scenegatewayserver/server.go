package main

import (
	"fmt"
	"net"

	"github.com/cpp2go/gonet"
	common "oldjon.com/common"
	fcmd "oldjon.com/fcmd"
)

type EchoTask struct {
	gonet.TcpTask
}

func NewEchoTask(conn net.Conn) *EchoTask {
	s := &EchoTask{
		TcpTask: *gonet.NewTcpTask(conn),
	}
	s.Derived = s
	return s
}

func (this *EchoTask) ParseMsg(data []byte, flag byte) bool {

	cmd := fcmd.UCmd(common.GetCmd(data))

	switch cmd {
	case fcmd.UCmd_Login:
		{
			req := &fcmd.ReqLogin{}
			ok := common.DecodeCmd(data, flag, req)
			if !ok {
				return false
			}

			fmt.Println("+++", cmd, ",", req.Account, ",", req.Password, ",")

			this.Verify()

			retCmd := &fcmd.RetLogin{
				UserId: 1,
			}
			this.SendCmd(fcmd.UCmd_Login, retCmd)
		}
	}

	return true
}

func (this *EchoTask) SendCmd(cmd fcmd.UCmd, msg common.Message) bool {
	data, flag, err := common.EncodeCmd(uint16(cmd), msg)
	if err != nil {
		fmt.Println("[服务] 发送失败 cmd:", cmd, ",len:", len(data), ",err:", err)
		return false
	}
	return this.AsyncSend(data, flag)
}

func (this *EchoTask) OnClose() {

}

type EchoServer struct {
	gonet.Service
	tcpser *gonet.TcpServer
}

var serverm *EchoServer

func EchoServer_GetMe() *EchoServer {
	if serverm == nil {
		serverm = &EchoServer{
			tcpser: &gonet.TcpServer{},
		}
		serverm.Derived = serverm
	}
	return serverm
}

func (this *EchoServer) Init() bool {
	err := this.tcpser.Bind(":80")
	if err != nil {
		fmt.Println("绑定端口失败")
		return false
	}
	return true
}

func (this *EchoServer) Reload() {

}

func (this *EchoServer) MainLoop() {
	conn, err := this.tcpser.Accept()
	if err != nil {
		return
	}
	NewEchoTask(conn).Start()
}

func (this *EchoServer) Final() bool {
	this.tcpser.Close()
	return true
}

func main() {

	EchoServer_GetMe().Main()

}
