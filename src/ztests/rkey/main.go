package main

import "fmt"

func main() {
	fmt.Println(RKey(1))
}

func RKey(ids ...interface{}) string {
	if len(ids) == 0 {
		return ""
	}
	var tmp []interface{}
	for i := 0; i < len(ids)-1; i++ {
		tmp = append(tmp, ids[i], ":")
	}
	tmp = append(tmp, ids[len(ids)-1])
	return fmt.Sprint(tmp...)
}
