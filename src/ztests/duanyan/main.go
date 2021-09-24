package main

import (
	"fmt"
	"time"
)

type AI interface {
	FA()
}

type BI interface {
	FB()
}

type AStruct struct {
}

func (this *AStruct) FA() {
	fmt.Println("FA")
}

func (this *AStruct) FB() {
	fmt.Println("FB")
}

func main() {
	//a := &AStruct{}
	//test(a)
	testbench()
}

func test(in interface{}) {
	b, ok := in.(BI)
	if !ok || b == nil {
		fmt.Println("duanyan err", ok)
		return
	}
	b.FB()
	a, ok := in.(AI)
	if !ok || a == nil {
		fmt.Println("duanyan err", ok)
		return
	}
	a.FA()
	fmt.Println("default")
}

type FAI interface {
	FA()
}

type Small struct {
	P int
}

func (this *Small) FA() {
	return
}

type Big struct {
	P [10000]int
}

func (this *Big) FA() {
	return
}

func testbench() { //指针对象的大小不影响断言效率
	var small interface{} = &Small{}
	b := &Big{}
	for i, _ := range b.P {
		b.P[i] = i
	}
	var big interface{} = b
	_ = big
	nowunix := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		a, _ := small.(FAI)
		a.FA()
	}
	mid := time.Now().UnixNano()
	fmt.Println(mid - nowunix)
	for i := 0; i < 100000000; i++ {
		/*b, ok := big.(FAI)
		if !ok {
			fmt.Println("err")
			return
		}*/
		b.FA()
	}
	end := time.Now().UnixNano()
	fmt.Println(end - mid)
}
