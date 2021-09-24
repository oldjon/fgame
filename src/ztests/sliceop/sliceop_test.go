//package main
//
//import (
//	"fmt"
//	"testing"
//	"time"
//)
//
//var count int
//
//func Test_SliceOp(b *testing.T) {
//	sli := []int64{}
//	go func() {
//		for {
//			if len(sli) > 0 {
//				fmtslice(sli)
//			}
//			sli = sli[:0]
//		}
//	}()
//	for i := int64(0); i < 1000; i++ {
//		sli = append(sli, i)
//	}
//	time.Sleep(time.Second * 10)
//	fmt.Println(count)
//}
//
//func fmtslice(s []int64) {
//	for i := range s {
//		_ = s[i]
//		count++
//	}
//}
