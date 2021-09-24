package com

const (
	MsgFlag_Compress uint8 = 1
	MsgFlag_UId      uint8 = 1 << 1
	MsgFlag_Err      uint8 = 1 << 2
	MsgFlag_Async    uint8 = 1 << 3
	MsgFlag_Push     uint8 = 1 << 4
	MsgFlag_AES      uint8 = 1 << 5
)
