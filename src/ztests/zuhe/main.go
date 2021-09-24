package main

import "fmt"

type Base struct {
}

func (this *Base) Hello() {
	fmt.Println("this is base")
	return
}

type A struct {
	Base
}

type B struct {
	Base
}

func main() {
	var b B
	b.Hello()
}
