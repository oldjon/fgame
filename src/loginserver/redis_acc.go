package main

import (
	"strconv"
	"strings"
	"time"

	"oldjon.com/com"
	"oldjon.com/dbobj"
	"oldjon.com/glog"
	"oldjon.com/redis-go-cluster"
)

func (this *RedisMgr) GetAccData(account string) (*db.Account, bool) {
	cluster := this.GetCacheCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		return nil, false
	}

	key := db.RKey(db.RKey_Account, account)
	data := &db.Account{}
	err := cluster.GetObject(key, data)
	if err != nil {
		return nil, false
	}
	return data, true
}

func (this *RedisMgr) GetUserIdByTel(tel string) (uint64, bool) {
	cluster := this.GetAccCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		return 0, false
	}
	key := db.RKey(db.RKey_Tel, tel)
	userid, err := redis.Uint64(cluster.Do("GET", key))
	if err != nil {
		glog.Error("[手机] 通过手机获取玩家id失败 ", tel)
		return 0, false
	}
	return userid, true
}

func (this *RedisMgr) GetAccount(userid uint64) (string, bool) {
	cluster := this.GetAccCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		return "", false
	}
	key := db.RKey(db.RKey_User, userid)
	account, err := redis.String(cluster.Do("HGET", key))
	if err != nil {
		glog.Error("[手机] 通过手机获取玩家id失败 ", userid)
		return "", false
	}
	return account, true
}

func (this *RedisMgr) GetByDevice(device string, platform uint32, clientaddr, phonetype string) (userid uint64, account string, isnew bool, errcode uint32) {
	cluster := this.GetAccCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		errcode = com.Err_DB
		return
	}
	key := db.RKey(db.RKey_Device, device)
	var err error
	if !cluster.Exists(key) {
		tmpuid, err := cluster.Incr("next.userid")
		if err != nil {
			glog.Error("[建号] 数据库失败 ", device, ",", err)
			errcode = com.Err_DB
			return
		}
		userid = 10000 + uint64(tmpuid)
		account = "FG" + strconv.Itoa(int(userid)+100000000)
		password := com.EncMd5(strconv.Itoa(int(userid) + int(time.Now().Unix())))
		if !cluster.HSet(key, "userid", userid) {
			glog.Error("[建号] 设置设备信息失败 ", device)
			errcode = com.Err_DB
			return
		}
		accdata := &db.Account{
			Id:       userid,
			Password: password,
			IsBind:   0,
		}
		if !cluster.SetObject(db.RKey(db.RKey_Account, strings.ToLower(account)), accdata) {
			glog.Error("[建号] 设置账号信息失败 ", device)
			errcode = com.Err_DB
			return
		}

		gamedata := &db.GameData{
			Account:  account,
			Device:   device,
			HeadIcon: com.RandBetweenUint32(0, 9), //默认头像
			HeadUrl:  "",                          //头像url
			RegTime:  time.Now().Unix(),
			Platform: platform, //平台
		}
		if !cluster.SetObject(db.RKey(db.RKey_GameData, userid), gamedata) {
			glog.Error("[建号] 设置游戏数据失败 ", device)
			errcode = com.Err_DB
			return
		}
		isnew = true
		glog.Info("[账号] 新号注册成功 ", userid, platform, device, clientaddr, phonetype)
	} else {
		userid, err = redis.Uint64(cluster.Do("HGET", key, "id"))
		if err != nil {
			glog.Error("[登录] 根据设备号获取玩家id失败 ", device)
			errcode = com.Err_DB
			return
		}
		var gamedata = &db.GameData{}
		err = cluster.GetObject(db.RKey(db.RKey_GameData, userid), gamedata)
		if err != nil {
			glog.Error("[登录] 获取玩家游戏数据失败 ", device)
			errcode = com.Err_DB
			return
		}
		account = gamedata.Account
	}
	return
}
