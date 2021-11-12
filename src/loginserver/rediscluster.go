package main

import (
	"oldjon.com/com"
)

type RedisMgr struct {
	com.RedisMgr
}

var redismsgr *RedisMgr

func RedisMgr_GetMe() *RedisMgr {
	if redismsgr == nil {
		redismsgr = &RedisMgr{}
	}
	return redismsgr
}
