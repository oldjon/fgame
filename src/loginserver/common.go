package main

import (
	"encoding/binary"
	"io/ioutil"
	"net/http"

	"oldjon.com/com"
	"oldjon.com/fcmd"
	"oldjon.com/glog"
	pb "oldjon.com/protobuf/proto"
)

func getRemoteAddr(req *http.Request) string {
	if req == nil {
		return ""
	}
	ip := req.Header.Get("fgame-ip")
	if ip != "" {
		return ip
	}
	return req.RemoteAddr
}

func GetPostWalues(req *http.Request) (service fcmd.Service, cmd fcmd.UCmd, userid uint64, data []byte, addr string, ok bool) {
	addr = getRemoteAddr(req)
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		glog.Error("[登录] 接收数据失败 ", addr)
		return
	}
	if len(data) < 12 {
		glog.Error("[登录] 请求数据长度不足 ", addr, ",", len(data))
		return
	}
	service = fcmd.Service(data[0])
	cmd = fcmd.UCmd(uint32(data[1]) | uint32(data[2])<<8 | uint32(data[3])<<16) //小端，低地址位存放数值低位
	userid = binary.LittleEndian.Uint64(data[4:])                               //小端，低地址位存放数值低位
	data = data[8:]
	ok = true
	return
}

func RetErrMsg(res http.ResponseWriter, errcode uint32) {
	buf := make([]byte, 5)
	buf[0] = com.MsgFlag_Err
	binary.LittleEndian.PutUint32(buf[1:], errcode)
	res.Write(buf)
}

func SendMsg(res http.ResponseWriter, message pb.Message) bool {
	data, err := pb.Marshal(message)
	if err != nil {
		return false
	}
	buf := make([]byte, 1+len(data))
	copy(buf[1:], data)
	res.Write(buf)
	return true
}
