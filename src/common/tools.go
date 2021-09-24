package com

import (
	"math/rand"
	"strconv"
	"time"

	"oldjon.com/glog"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandBetween(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return min + rand.Int63n(max-min+1)
}

func RandBetweenUint32(min, max uint32) uint32 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return min + uint32(rand.Int63n(int64(max-min+1)))
}

func RandBetweenFloat(min, max float64) float64 {
	if min == max {
		return min
	}
	if min > max {
		min, max = max, min
	}
	return min + rand.Float64()*(max-min)
}

func StrToUint32(s string) uint32 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		glog.ErrorDepth(1, "[字符串转换] 转换失败 ", s, ",", err)
		return 0
	}
	return uint32(i)
}

func StrToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		glog.ErrorDepth(1, "[字符串转换] 转换失败 ", s, ",", err)
		return 0
	}
	return i
}

func StrToInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		glog.ErrorDepth(1, "[字符串转换] 转换失败 ", s, ",", err)
		return 0
	}
	return int32(i)
}

func StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		glog.ErrorDepth(1, "[字符串转换] 转换失败 ", s, ",", err)
		return 0
	}
	return i
}

func StrToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		glog.ErrorDepth(1, "[字符串转换] 转换失败 ", s, ",", err)
		return 0
	}
	return int(i)
}
