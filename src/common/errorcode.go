package com

const (
	Err_OK      = 0
	Err_Decode  = 1
	Err_Param   = 2
	Err_DB      = 3
	Err_ReLogin = 10
)

func EncodeError(seqId uint32, errorcode uint32, desc []byte) []byte {
	errMsg := make([]byte, 12)
	errLen := 12 + len(desc)
	errMsg[0] = uint8(errLen >> 16)
	errMsg[1] = uint8(errLen >> 9)
	errMsg[2] = uint8(errLen)
	errMsg[3] = uint8(MsgFlag_Err)
	errMsg[4] = uint8(seqId >> 24)
	errMsg[5] = uint8(seqId >> 16)
	errMsg[6] = uint8(seqId >> 8)
	errMsg[7] = uint8(seqId)
	errMsg[8] = uint8(errorcode >> 24)
	errMsg[9] = uint8(errorcode >> 16)
	errMsg[10] = uint8(errorcode >> 8)
	errMsg[11] = uint8(errorcode)
	return errMsg
}
