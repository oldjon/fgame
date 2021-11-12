package com

import (
	"strings"
	"time"

	"oldjon.com/glog"
	"oldjon.com/redis-go-cluster"
)

type RedisMgr struct {
	cacheCluster *redis.Cluster
	accCluster   *redis.Cluster
	lockCluster  *redis.Cluster
}

func (this *RedisMgr) Init(clusters map[string]string) bool {
	var ok bool
	for c, addr := range clusters {
		switch c {
		case "cache":
			this.cacheCluster, ok = this.NewRedisCluster(addr)
			if !ok {
				return false
			}
		case "acc":
			this.accCluster, ok = this.NewRedisCluster(addr)
			if !ok {
				return false
			}
		case "lock":
			this.lockCluster, ok = this.NewRedisCluster(addr)
			if !ok {
				return false
			}
		default:
			glog.Error("[redis] 未支持的集群类型 ", c, ",", addr)
			return false
		}
	}
	return true
}

func (this *RedisMgr) NewRedisCluster(addrs string) (*redis.Cluster, bool) {
	if len(addrs) == 0 {
		glog.Error("[启动] redis集群配置错误 ", addrs)
		return nil, false
	}
	addrlist := strings.Split(addrs, "|")
	if len(addrlist) < 1 {
		glog.Error("[启动] redis集群配置错误 ", addrs)
		return nil, false
	}
	cacheCluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   addrlist,
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		glog.Error("[启动] redis.New error: %s", err.Error(), ",", addrs)
		return nil, false
	}
	return cacheCluster, true
}

func (this *RedisMgr) GetCacheCluster() *redis.Cluster {
	return this.cacheCluster
}

func (this *RedisMgr) GetAccCluster() *redis.Cluster {
	return this.accCluster
}

func (this *RedisMgr) GetLockCluster() *redis.Cluster {
	return this.lockCluster
}

func (this *RedisMgr) Close() {
	if this.cacheCluster != nil {
		this.cacheCluster.Close()
	}
	if this.accCluster != nil {
		this.accCluster.Close()
	}
	if this.lockCluster != nil {
		this.lockCluster.Close()
	}
	return
}
