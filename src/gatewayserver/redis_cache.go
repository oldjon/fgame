package main

import (
	"oldjon.com/glog"
)

func (this *RedisMgr) LoadObject(key string, object interface{}) bool {
	cluster := this.GetCacheCluster()
	if cluster == nil {
		glog.Error("[Redis] 获取cluster失败 ")
		return false
	}
	if !cluster.Exists(key) {
		return false
	}

	err := cluster.GetObject(key, object)
	if err != nil {
		return false
	}

	if !cluster.Del(key) {
		return false
	}
	return true
}
