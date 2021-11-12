package main

import (
	"oldjon.com/glog"
)

func (this *RedisMgr) SetToken(key string, object interface{}, expritetime int64) bool {
	cluster := this.GetCacheCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		return false
	}

	if !cluster.SetObject(key, object) {
		return false
	}

	if !cluster.Expire(key, expritetime) {
		return false
	}
	return true
}
