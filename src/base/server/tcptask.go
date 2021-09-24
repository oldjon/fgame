package server

import (
	"io"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"oldjon.com/glog"
)

type ITcpTask interface {
	ParseMsg(data []byte) bool
	OnClose()
}

const (
	cmd_max_size    = 128 * 1024
	cmd_header_size = 4 //3字节指令，1字节flag
	cmd_verify_time = 30
)

type TCPTask struct {
	closed     int32
	verified   bool
	stopedChan chan struct{}
	recvBuff   *ByteBuffer // 考虑使用ringbuff
	sendBuff   *ByteBuffer // 考虑使用ringbuff
	sendMutex  sync.Mutex
	sendChan   chan struct{}
	Conn       net.Conn
	Derived    ITcpTask
}

func NewTCPTask(conn net.Conn) *TCPTask {
	return &TCPTask{
		closed:     -1,
		verified:   false,
		Conn:       conn,
		stopedChan: make(chan struct{}, 1),
		recvBuff:   NewByteBuffer(),
		sendBuff:   NewByteBuffer(),
		sendChan:   make(chan struct{}, 1),
	}
}

func (this *TCPTask) SendSignal() {
	select {
	case this.sendChan <- struct{}{}:
	default:
	}
	return
}

func (this *TCPTask) RemoteAddr() string {
	if this.Conn == nil {
		return ""
	}
	return this.Conn.RemoteAddr().String()
}

func (this *TCPTask) LocalAddr() string {
	if this.Conn == nil {
		return ""
	}
	return this.Conn.LocalAddr().String()
}

func (this *TCPTask) IsClosed() bool {
	return atomic.LoadInt32(&this.closed) != 0
}

func (this *TCPTask) Stop() bool {
	if this.IsClosed() {
		glog.Error("[连接] 关闭失败 ", this.RemoteAddr())
		return false
	}
	select {
	case this.stopedChan <- struct{}{}:
	default:
		glog.Error("[连接] 关闭失败 ", this.RemoteAddr())
		return false
	}
	return true
}

func (this *TCPTask) Start() {
	if !atomic.CompareAndSwapInt32(&this.closed, -1, 0) {
		return
	}
	job := &sync.WaitGroup{}
	job.Add(1)
	go this.SendLoop(job)
	go this.RecvLoop()
	job.Wait()
	glog.Info("[连接] 收到连接 ", this.RemoteAddr())
	return
}

func (this *TCPTask) Close() {
	if !atomic.CompareAndSwapInt32(&this.closed, 0, 1) {
		return
	}
	glog.Info("[连接] 断开连接 ", this.RemoteAddr())
	this.Conn.Close()
	this.recvBuff.Reset()
	this.sendBuff.Reset()
	select {
	case this.stopedChan <- struct{}{}:
	default:
		glog.Error("[连接] 关闭失败 ", this.RemoteAddr())
	}
	this.Derived.OnClose()
	return
}

func (this *TCPTask) Reset() bool {
	if atomic.LoadInt32(&this.closed) != 1 {
		return false
	}
	if !this.IsVerified() {
		return false
	}
	this.closed = -1
	this.verified = true
	this.stopedChan = make(chan struct{})
	glog.Info("[连接] 重置连接 ", this.RemoteAddr())
	return true
}

func (this *TCPTask) Verify() {
	this.verified = true
	return
}

func (this *TCPTask) IsVerified() bool {
	return this.verified
}

func (this *TCPTask) Terminate() {
	this.Close()
}

func (this *TCPTask) SendBytes(buffer []byte) bool {
	if this.IsClosed() {
		return false
	}
	this.sendMutex.Lock()
	this.sendBuff.Append(buffer...)
	this.sendMutex.Unlock()
	this.SendSignal()
	return true
}

func (this *TCPTask) readAtLeast(buff *ByteBuffer, neednum int) error {
	buff.WrGrow(neednum)
	n, err := io.ReadAtLeast(this.Conn, buff.WrBuf(), neednum)
	buff.WrFlip(n)
	return err
}

func (this *TCPTask) RecvLoop() {
	defer func() {
		this.Close()
		if err := recover(); err != nil {
			glog.Error("[异常] ", err, "\n", string(debug.Stack()))
		}
	}()

	var (
		neednum   int
		err       error
		totalsize int
		datasize  int
		msgbuff   []byte
	)

	for {

		totalsize = this.recvBuff.RdSize()
		if totalsize <= cmd_header_size {
			neednum = cmd_header_size - totalsize
			err = this.readAtLeast(this.recvBuff, neednum)
			if err != nil {
				glog.Error("[连接] 接收数据失败 ", this.RemoteAddr(), ",", err)
				return
			}
			totalsize = this.recvBuff.RdSize()
		}

		msgbuff = this.recvBuff.RdBuf()

		datasize = int(msgbuff[0]<<16) | int(msgbuff[1]<<8) | int(msgbuff[2])
		if datasize > cmd_max_size {
			glog.Error("[连接] 数据长度超过最大值 ", this.RemoteAddr(), ",", datasize)
			return
		} else if datasize < cmd_header_size {
			glog.Error("[连接] 数据长度不足最小值 ", this.RemoteAddr(), ",", datasize)
			return
		}

		if totalsize < datasize {
			neednum = datasize - totalsize
			err = this.readAtLeast(this.recvBuff, neednum)
			if err != nil {
				glog.Error("[连接] 接收数据失败 ", this.RemoteAddr(), ",", err)
				return
			}
			msgbuff = this.recvBuff.RdBuf()
		}

		this.Derived.ParseMsg(msgbuff[:datasize])
		this.recvBuff.RdFlip(datasize)
	}
}

func (this *TCPTask) SendLoop(job *sync.WaitGroup) {
	defer func() {
		this.Close()
		if err := recover(); err != nil {
			glog.Error("[异常] ", err, "\n", string(debug.Stack()))
		}
	}()

	var (
		tmpByte  = NewByteBuffer()
		timeout  = time.NewTimer(time.Second * cmd_verify_time)
		writenum int
		err      error
	)

	defer timeout.Stop()

	job.Done()

	for {
		select {
		case <-this.sendChan:
			for {
				this.sendMutex.Lock()
				if this.sendBuff.RdReady() {
					tmpByte.Append(this.sendBuff.RdBuf()[:this.sendBuff.RdSize()]...)
					this.sendBuff.Reset()
				}
				this.sendMutex.Unlock()

				if !tmpByte.RdReady() {
					break
				}

				writenum, err = this.Conn.Write(tmpByte.RdBuf()[:tmpByte.RdSize()])
				if err != nil {
					glog.Error("[连接] 发送失败 ", this.RemoteAddr(), ",", err)
					return
				}
				tmpByte.RdFlip(writenum)
			}
		case <-this.stopedChan:
			return
		case <-timeout.C:
			if !this.IsVerified() {
				glog.Error("[连接] 验证超时 ", this.RemoteAddr())
				return
			}
		}
	}
}
