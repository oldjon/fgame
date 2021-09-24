package server

import (
	"net"
	"time"

	"oldjon.com/glog"
)

type TCPServer struct {
	listener *net.TCPListener
}

func (this *TCPServer) Bind(address string) error {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		glog.Error("[服务] 解析失败 ", address)
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		glog.Error("[服务] 侦听失败 ", address)
		return err
	}

	glog.Info("[服务] 侦听成功 ", address)
	this.listener = listener
	return nil
}

func (this *TCPServer) Accept() (*net.TCPConn, error) {

	this.listener.SetDeadline(time.Now().Add(time.Second * 1))

	conn, err := this.listener.AcceptTCP()
	if err != nil {
		return nil, err
	}

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Minute)
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)

	return conn, nil
}

func (this *TCPServer) BindAccept(address string, handler func(conn *net.TCPConn)) error {
	err := this.Bind(address)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := this.Accept()
			if err != nil {
				continue
			}
			handler(conn)
		}
	}()
	return nil
}

func (this *TCPServer) CLose() error {
	return this.listener.Close()
}
