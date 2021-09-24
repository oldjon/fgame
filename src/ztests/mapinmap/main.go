package main

import "fmt"

func main() {
	mm := make(map[int]map[int]int)
	m, ok := mm[0]
	if !ok {
		m = make(map[int]int)
		mm[0] = m
	}
	m[0] = 1
	fmt.Println(mm)
}
