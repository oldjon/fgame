package redis

import (
	"strings"

	"oldjon.com/glog"
)

//********************KEY********************
func (this *Cluster) Del(key string) bool {
	n, err := Int(this.Do("DEL", key))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] DEL操作失败 ", err)
		return false
	}
	if n != 1 {
		return false
	}
	return true
}

func (this *Cluster) Exists(key string) bool {
	n, err := Int(this.Do("EXISTS", key))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] EXISTS操作失败 ", err)
		return false
	}
	if n != 1 {
		return false
	}
	return true
}

func (this *Cluster) Expire(key string, secs int64) bool {
	n, err := Int(this.Do("EXPIRE", key, secs))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] EXPIRE操作失败 ", err)
		return false
	}
	if n != 1 {
		return false
	}
	return true
}

func (this *Cluster) ExpireAt(key string, expiretimestamp int64) bool {
	n, err := Int(this.Do("EXPIREAT", key, expiretimestamp))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] EXPIREAT操作失败 ", err)
		return false
	}
	if n != 1 {
		return false
	}
	return true
}

func (this *Cluster) TTL(key string) int64 {
	n, err := Int64(this.Do("TTL", key))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] TTL操作失败 ", err)
		return -2
	}
	return n
}

//删除给定key的过期时间
func (this *Cluster) Persist(key string) bool {
	n, err := Int(this.Do("PERSIST", key))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] PERSIST操作失败 ", err)
		return false
	}
	if n != 1 {
		glog.ErrorDepth(1, "[Redis] PERSIST操作失败 ", err, n)
		return false
	}
	return true
}

func (this *Cluster) PExpire(key string, milsecs int64) bool {
	n, err := Int(this.Do("PEXPIRE", key, milsecs))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] PEXPIRE操作失败 ", err)
		return false
	}
	if n != 1 {
		glog.ErrorDepth(1, "[Redis] PEXPIRE操作失败 ", err, n)
		return false
	}
	return true
}

func (this *Cluster) PExpireAt(key string, expiretimestampmilsec int64) bool {
	n, err := Int(this.Do("PEXPIREAT", key, expiretimestampmilsec))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] PEXPIREAT操作失败 ", err)
		return false
	}
	if n != 1 {
		glog.ErrorDepth(1, "[Redis] PEXPIREAT操作失败 ", err, n)
		return false
	}
	return true
}

func (this *Cluster) PTTL(key string) int64 {
	n, err := Int64(this.Do("PTTL", key))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] PTTL操作失败 ", err)
		return -2
	}
	return n
}

//********************STRING********************
func (this *Cluster) Incr(key string) (int64, error) {
	return Int64(this.Do("INCR", key))
}

func (this *Cluster) Decr(key string) (int64, error) {
	return Int64(this.Do("DECR", key))
}

func (this *Cluster) Incrby(key string, num int64) (int64, error) {
	return Int64(this.Do("INCRBY", key, num))
}

func (this *Cluster) Decrby(key string, num int64) (int64, error) {
	return Int64(this.Do("DECRBY", key, num))
}

func (this *Cluster) Set(key string, val interface{}) bool {
	ok, err := String(this.Do("SET", key, val))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] SET操作失败 ", err)
		return false
	}
	if strings.ToUpper(ok) != "OK" {
		glog.ErrorDepth(1, "[Redis] SET操作失败 ", err, ",", ok)
		return false
	}
	return true
}

func (this *Cluster) SetEX(key string, val interface{}, ttl int64) bool {
	ok, err := String(this.Do("SETEX", key, ttl, val))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] SET操作失败 ", err)
		return false
	}
	if strings.ToUpper(ok) != "OK" {
		glog.ErrorDepth(1, "[Redis] SET操作失败 ", err, ",", ok)
		return false
	}
	return true
}

func (this *Cluster) SetNX(key string, val interface{}) bool {
	n, err := Int(this.Do("SETNX", key, val))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] SETNX操作失败 ", err)
		return false
	}
	if n != 1 {
		glog.ErrorDepth(1, "[Redis] SETNX操作失败 ", err, ",", n)
		return false
	}
	return true
}

func (this *Cluster) GetObject(key string, obj interface{}) error {
	fields, err := GetStructFields(obj)
	if err != nil {
		return err
	}
	v, err := Values(this.Do("HMGET", Args{}.Add(key).AddFlat(fields)...))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] HMGET操作失败 ", err)
		return err
	}
	return ScanStructByValue(v, fields, obj)
}

func (this *Cluster) SetObject(key string, obj interface{}) bool {
	ok, err := String(this.Do("HMSET", Args{}.Add(key).AddFlat(obj)...))
	if err != nil {
		glog.ErrorDepth(1, "[Redis] HMSET操作失败 ", err)
		return false
	}
	if strings.ToUpper(ok) != "OK" {
		glog.ErrorDepth(1, "[Redis] HMSET操作失败 ", err, ",", ok)
		return false
	}
	return true
}
