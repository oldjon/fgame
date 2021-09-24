package com

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"

	"oldjon.com/fcmd"
	"oldjon.com/glog"
	"oldjon.com/protobuf/proto"
)

const (
	MaxCompressSize = 1024
	CmdHeaderSize   = 8
	CmdSeqSize      = 4
)

//async msg | 3bytes length | 1byte flag | 1byte service | 3bytes cmd     | 		data		|
//sync msg  | 3bytes length | 1byte flag | 1byte service | 3bytes cmd     | 4bytes seqid | data |
//err msg   | 3bytes length | 1byte flag | 4bytes seqid  | 2bytes errcode |

type Message = proto.Message

/*type Message interface {
	Marshal() (data []byte, err error)
	MarshalTo(data []byte) (n int, err error)
	Size() (n int)
	Unmarshal(data []byte) error
}*/

func zlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, err := w.Write(src)
	if err != nil {
		return nil
	}
	w.Close()
	return in.Bytes()
}

func zlibUnCompress(src []byte) []byte {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return nil
	}
	return out.Bytes()
}

// 生成二进制数据,返回数据和是否压缩标识
func EncodeCmd(cmd uint16, msg Message) ([]byte, byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		glog.Error("[协议] 编码错误 ", err)
		return nil, 0, err
	}
	var (
		mflag byte
		mbuff []byte
	)
	mflag = 0 //TODO
	if len(data) >= MaxCompressSize {
		mbuff = zlibCompress(data)
		mflag |= MsgFlag_Compress
	} else {
		mbuff = data
	}
	p := make([]byte, len(mbuff)+CmdHeaderSize)
	binary.LittleEndian.PutUint16(p[0:], cmd)
	copy(p[CmdHeaderSize:], mbuff)
	return p, mflag, nil
}

// 生成protobuf数据
func DecodeCmd(buf []byte, pb Message) (userid uint64, ok bool) {
	if len(buf) < CmdHeaderSize {
		glog.Error("[协议] 数据错误 ", buf)
		return 0, false
	}
	var start = CmdHeaderSize
	flag := buf[3]
	if flag&MsgFlag_UId == MsgFlag_UId {
		uid, size := DecodeVarintReverse(buf, len(buf)-1)
		buf = buf[:size]
		userid = uid
	}

	if flag&MsgFlag_Async != MsgFlag_Async {
		if len(buf) < CmdHeaderSize+CmdSeqSize {
			glog.Error("[协议] 数据错误 ", buf)
			return
		}
		start += CmdSeqSize
	}

	var mbuff []byte
	if flag&MsgFlag_Compress == MsgFlag_Compress {
		mbuff = zlibUnCompress(buf[start:]) //后续通过接口来来处理
	} else {
		mbuff = buf[start:]
	}
	err := proto.Unmarshal(mbuff, pb)
	if err != nil {
		glog.Error("[协议] 解码错误 ", err, ",", mbuff)
		return 0, false
	}
	return 0, true
}

func VarintSize(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func EncodeVarintReverse(data []byte, offset int, v uint64) int {
	for v >= 0x80 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset--
	}
	data[offset] = uint8(v)
	return offset - 1
}

func DecodeVarintReverse(data []byte, offset int) (p uint64, newOffset int) {
	newOffset = offset
	for shift := uint(0); newOffset >= 0; shift += 7 {
		b := data[newOffset]
		p |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			break
		}
		newOffset--
	}
	return
}

// 获取服务
func GetService(buf []byte) fcmd.Service {
	if len(buf) < CmdHeaderSize {
		return 0
	}
	return fcmd.Service(buf[4])
}

// 获取指令号
func GetUCmd(buf []byte) (ucmd fcmd.UCmd) {
	if len(buf) < CmdHeaderSize {
		return
	}
	ucmd = fcmd.UCmd(uint32(buf[5])<<16 | uint32(buf[6])<<8 | uint32(buf[7]))
	return
}

func GetSeqId(buf []byte) (seqid uint32) {
	if len(buf) < CmdHeaderSize {
		return
	}
	if buf[3]&MsgFlag_Async == 0 {
		if len(buf) < CmdHeaderSize+CmdSeqSize {
			return
		}
		seqid = uint32(buf[8])<<24 | uint32(buf[9])<<16 | uint32(buf[10])<<8 | uint32(buf[11])
	}
	return
}

func MarshalUserId(userid uint64, data []byte) []byte {
	l := len(data) + VarintSize(userid)
	buff := make([]byte, l)
	copy(buff[0:], data)
	buff[0] = uint8(l >> 16)
	buff[1] = uint8(l >> 8)
	buff[2] = uint8(l)
	buff[3] |= MsgFlag_UId
	EncodeVarintReverse(buff, l-1, userid)
	binary.LittleEndian.PutUint64(buff[len(data):], userid)
	return buff
}

func UnmarshalUserId(data []byte) (userid uint64, out []byte, ok bool) {
	if data[3]&MsgFlag_UId != MsgFlag_UId {
		return 0, nil, false
	}
	l := len(data)
	var size int
	userid, size = DecodeVarintReverse(data, l-1)
	l = size
	data[0] = uint8(l >> 16)
	data[1] = uint8(l >> 8)
	data[2] = uint8(l)
	data[3] &= ^MsgFlag_UId
	return userid, data[:l], true
}

//async msg | 3bytes length | 1byte flag | 1byte service | 3bytes cmd     | 		data		|
func CreateServiceMsg(service fcmd.Service, ucmd fcmd.UCmd, flag uint8, pbdata []byte) []byte {
	l := len(pbdata) + 8
	buff := make([]byte, l)
	buff[0] = uint8(l >> 16)
	buff[1] = uint8(l >> 8)
	buff[2] = uint8(l)
	buff[3] = flag
	buff[4] = uint8(service)
	buff[5] = uint8(ucmd >> 16)
	buff[6] = uint8(ucmd >> 8)
	buff[7] = uint8(ucmd)
	copy(buff[8:], pbdata)
	return buff
}
