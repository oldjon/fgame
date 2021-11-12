package com

const CaptchaLength = 4

const (
	State_Ok uint32 = iota
)

const (
	LoginType_Device          = 1 //设备登录
	LoginType_AccountPassword = 2 //账户密码登录
	LoginType_TelPassword     = 3 //手机号密码登录
	LoginType_TelCaptcha      = 4 //手机号验证码登录
)
