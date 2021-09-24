package main

import (
	"fmt"
	"time"
)

func main() {
	A()
}

var count int

func A() {
	sli := make([]int64, 0, 10000)
	go func() {
		for {
			if len(sli) > 0 {
				fmtslice(sli)
			}
			sli = sli[:0]
		}
	}()
	for i := int64(0); i < 2000; i++ {
		sli = append(sli, i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(count)
}

func fmtslice(s []int64) {
	for range s {
		//_ = s[i]
		count++
	}
}
