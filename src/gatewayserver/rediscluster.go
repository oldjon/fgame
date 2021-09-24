package main

import (
	"strings"
	"time"

	"oldjon.com/env"
	"oldjon.com/glog"
	"oldjon.com/redis-go-cluster"
)

type RedisMgr struct {
	cacheCluster *redis.Cluster
}

var redismsgr *RedisMgr

func RedisMgr_GetMe() *RedisMgr {
	if redismsgr == nil {
		redismsgr = &RedisMgr{}
	}
	return redismsgr
}

func (this *RedisMgr) Init() bool {
	redisclusterstr := env.Get("gatewayserver", "rediscluster")
	if len(redisclusterstr) == 0 {
		glog.Error("[启动] redis集群配置错误 ", redisclusterstr)
		return false
	}
	redisclusteraddr := strings.Split(redisclusterstr, "|")
	if len(redisclusteraddr) < 1 {
		glog.Error("[启动] redis集群配置错误 ", redisclusterstr)
		return false
	}
	cacheCluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   redisclusteraddr,
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})

	if err != nil {
		glog.Error("[启动] redis.New error: %s", err.Error())
		return false
	}

	this.cacheCluster = cacheCluster
	return true
}

func (this *RedisMgr) GetCacheCluster() *redis.Cluster {
	return this.cacheCluster
}
