package com

import (
	"time"

	"oldjon.com/glog"
)

type HandlerFunc func() uint32

func (this *RedisMgr) Lock(key string) bool {
	cluster := this.GetLockCluster()
	if cluster == nil {
		glog.Error("[分布式锁] 获取cluster失败 LockCluster")
		return false
	}
	i := 0
	redlock := "rl:" + key
	for ; i < 25; i++ {
		ret, _ := cluster.Do("SET", redlock, 1, "PX", 2000, "NX")
		if ret != nil {
			return true
		}
		time.Sleep(time.Millisecond * time.Duration(i+1) * 5)
	}
	if i >= 25 {
		glog.Info("[分布式锁] 加锁超时 ", key)
	}
	return false
}

func (this *RedisMgr) Unlock(key string) {
	cluster := this.GetLockCluster()
	if cluster == nil {
		glog.Error("[分布式锁] 获取cluster失败 LockCluster")
		return
	}
	cluster.Del("rl:" + key)
	return
}

func (this *RedisMgr) Safely(key string, hanler HandlerFunc) (uint32, bool) {
	var (
		now      = time.Now()
		funcmsec int64
	)
	defer func() {
		usemsec := time.Now().Sub(now).Milliseconds()
		if usemsec > 1000 {
			glog.Warning("[卡顿] 执行时间 ", key, ",", usemsec, ",", funcmsec)
		}
	}()
	if !this.Lock(key) {
		return 0, false
	}
	funcmstart := time.Now()
	retcode := hanler()
	funcmsec = time.Now().Sub(funcmstart).Milliseconds()
	this.Unlock(key)
	return retcode, true
}
