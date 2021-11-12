package main

import (
	"net/http"
	"strconv"
	"time"

	"oldjon.com/com"
	"oldjon.com/dbobj"
	"oldjon.com/fcmd"
	"oldjon.com/glog"
	pb "oldjon.com/protobuf/proto"
)

// /time
func HandleTime(res http.ResponseWriter, req *http.Request) {
	nowunix := time.Now().Unix()
	res.Write([]byte(strconv.Itoa(int(nowunix))))
}

// /login
func HandleLogin(res http.ResponseWriter, req *http.Request) {
	_, cmd, userid, data, clientaddr, ok := GetPostWalues(req)
	if !ok {
		return
	}
	glog.Info("[登录] 收到请求 ", clientaddr, ",", cmd, ",", len(data))
	switch cmd {
	case fcmd.UCmd_Login:
		reqmsg := &fcmd.ReqLogin{}
		err := pb.Unmarshal(data, reqmsg)
		if err != nil {
			RetErrMsg(res, com.Err_Param)
			glog.Error("[登录] 消息反序列化失败 ", clientaddr)
			return
		}
		if reqmsg.Account == "" && reqmsg.Device == "" {
			RetErrMsg(res, com.Err_Param)
			glog.Error("[登录] 参数错误 ", clientaddr)
			return
		}
		if errcode := CheckVersion(reqmsg.Version, reqmsg.ResVersion, reqmsg.APKVersion); errcode > 0 {
			RetErrMsg(res, errcode)
			glog.Error("[登录] 版本过低 ", clientaddr, ",", errcode, ",", reqmsg.Version, ",", reqmsg.ResVersion, ",", reqmsg.APKVersion)
			return
		}
		var (
			logintype int
			isbind    bool
			account   string
			isnew     bool
			tel       string
		)

		if com.StrToUint64(reqmsg.Account) > 0 { //纯数字认为是手机号
			tel = reqmsg.Account
			reqmsg.Account = ""
		}

		if reqmsg.Account != "" {
			//账号密码登录
			if reqmsg.Password == "" {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 密码为空 ", clientaddr, ",", reqmsg.Account)
				return
			}
			acc, ok := RedisMgr_GetMe().GetAccData(reqmsg.Account)
			if !ok {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 获取账号失败 ", clientaddr, ",", reqmsg.Account)
				return
			}
			if acc.Password != reqmsg.Password {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 密码错误 ", clientaddr, ",", reqmsg.Account)
				return
			}
			userid = acc.Id
			isbind = acc.IsBind == 1
			account = reqmsg.Account
			logintype = com.LoginType_AccountPassword
		} else if tel != "" {
			//手机密码登录
			if reqmsg.Password == "" {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 密码为空 ", clientaddr, ",", reqmsg.Account)
				return
			}
			userid, ok = RedisMgr_GetMe().GetUserIdByTel(tel)
			if !ok {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 手机号不存在 ", clientaddr, ",", tel)
				return
			}
			account, ok = RedisMgr_GetMe().GetAccount(userid)
			if !ok {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 账号不存在 ", clientaddr, ",", tel)
				return
			}

			acc, ok := RedisMgr_GetMe().GetAccData(account)
			if !ok {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 获取账号失败 ", clientaddr, ",", reqmsg.Account)
				return
			}
			if acc.Password != reqmsg.Password {
				RetErrMsg(res, com.Err_AccOrPassword)
				glog.Error("[登录] 密码错误 ", clientaddr, ",", reqmsg.Account)
				return
			}
			isbind = acc.IsBind == 1
			logintype = com.LoginType_TelPassword
		} else if reqmsg.Tel != "" {
			if len(reqmsg.Tel) != 11 {
				RetErrMsg(res, com.Err_Tel)
				glog.Error("[登录] 手机号错误 ", clientaddr, ",", reqmsg.Tel)
				return
			}
			if len(reqmsg.Captcha) != com.CaptchaLength {
				RetErrMsg(res, com.Err_Captcha)
				glog.Error("[登录] 验证码错误 ", clientaddr, ",", reqmsg.Tel, ",", reqmsg.Captcha)
				return
			}
			//TODO 验证码校验逻辑
		} else if reqmsg.Device != "" { //设备号登录
			var errcode uint32
			userid, account, isnew, errcode = RedisMgr_GetMe().GetByDevice(reqmsg.Device, reqmsg.Platform, clientaddr, reqmsg.PhoneType)
			if errcode != com.Err_OK {
				glog.Error("[登录] 设备号登录失败 ", reqmsg.Device, ",", reqmsg.Platform, ",", reqmsg.APKVersion, ",",
					reqmsg.SystemOS, ",", reqmsg.PhoneType, ",", clientaddr, ",", reqmsg.Net, ",", reqmsg.Channel)
				RetErrMsg(res, errcode)
				return
			}
			logintype = com.LoginType_Device
		} else {
			RetErrMsg(res, com.Err_AccOrPassword)
			glog.Error("[登录] 密码为空 ", clientaddr, ",", reqmsg.Account)
			return
		}

		gateaddr := LoginServer_GetMe().GetGatewayServer(userid) //暂时不考虑按三网分配
		if gateaddr == "" {
			RetErrMsg(res, com.Err_Service)
			return
		}
		tokenkey := com.GenerateKey(userid)
		token := db.TokenData{
			Id:       userid,
			Account:  account,
			Platform: reqmsg.Platform,
		}
		if !RedisMgr_GetMe().SetToken(tokenkey, token, 120) {
			RetErrMsg(res, com.Err_DB)
			return
		}

		retmsg := &fcmd.RetLogin{
			UserId:       userid,
			Account:      account,
			IsBind:       isbind,
			GatewayAddr:  gateaddr,
			GatewayToken: tokenkey,
		}
		SendMsg(res, retmsg)

		glog.Info("[登录] 初始化成功 ", userid, ",", account, ",", reqmsg.Platform, ",", isnew, ",", reqmsg.APKVersion, ",", reqmsg.SystemOS, ",",
			reqmsg.Device, ",", reqmsg.PhoneType, ",", clientaddr, ",", reqmsg.Channel, ",", reqmsg.Net, ",", gateaddr, ",", logintype)
	//TODO 获取验证码等逻辑
	default:
		glog.Error("[登录] 未知指令 ", cmd)
		RetErrMsg(res, com.Err_Cmd)
		return
	}
}

func CheckVersion(_, _, _ string) uint32 { //TODO
	return com.Err_OK
}
