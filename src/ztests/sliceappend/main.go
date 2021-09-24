package main

func main() {
	l := make([]int, 10)
	for i := 0; i < len(l); {
		x := l[i]
		if x == 0 {
			l = append(l[:i], l[i+1:]...)
		} else {
			i++
		}
	}
}
