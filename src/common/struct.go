package com

type TokenData struct {
	Id       uint64 `redis:"Id"`
	Account  string `redis:"Account"`
	RegDev   string `redis:"RegDev"`
	Platform uint32 `redis:"PlatForm"`
}
