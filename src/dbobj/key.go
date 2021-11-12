package db

import (
	"fmt"
)

const (
	RKey_Account    = 1 //账户信息
	RKey_User       = 2 //玩家信息
	RKey_Tel        = 3 //手机号信息
	RKey_Device     = 4 //设备号
	RKey_DeviceList = 5 //设备列表
	RKey_GameData   = 6 //游戏数据
)

//构建redis key
func RKey(ids ...interface{}) string {
	if len(ids) == 0 {
		return ""
	}
	var tmp []interface{}
	for i := 0; i < len(ids)-1; i++ {
		tmp = append(tmp, ids[i], ":")
	}
	tmp = append(tmp, ids[len(ids)-1])
	return fmt.Sprint(tmp...)
}
