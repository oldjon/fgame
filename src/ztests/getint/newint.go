package main

import (
	"fmt"
	"sort"
)

func main() {
	var n = 8123
	fmt.Println(getint(n))
}

func getint(n int) int {
	if n <= 0 {
		return 0
	}
	m := []int{}
	var pre int
	var l uint
	for n > 0 {
		//m[n%10] = m[n%10] + 1
		m = append(m, n%10)
		pre = n % 10
		l++
		n = n / 10
	}

	pn := count(m, pre)
	if pn > 1 {
		m = removebynum(m, pre, pn-1)
		var f int
		for i := pn; i > 1; i-- {
			f += s10(pre, l)
			l--
		}
		fmt.Println(f)
		a := getmax(m, pre)
		if a == 0 {
			return 0
		}
		m = remove(m, f)
		f += s10(a, l)
		l--
		sort.Slice(m, func(i, j int) bool {
			return m[i] > m[j]
		})
		for _, v := range m {
			f = f + s10(v, l)
			l--
			fmt.Println(f)
		}
		return f
	} else {
		f := getmax(m, pre)
		if f == 0 {
			return 0
		}
		m = remove(m, f)
		sort.Slice(m, func(i, j int) bool {
			return m[i] > m[j]
		})

		f = s10(f, l)
		l--
		for _, v := range m {
			f = f + s10(v, l)
			l--
			fmt.Println(f)
		}
		return f
	}
	return 0
}

func count(m []int, p int) (n int) {
	for _, v := range m {
		if v == p {
			n++
		}
	}
	return n
}

func remove(m []int, p int) []int {
	for i, v := range m {
		if v == p {
			m = append(m[:i], m[i+1:]...)
			return m
		}
	}
	return nil
}

func removebynum(m []int, p, n int) []int {
	var ret []int
	for i, v := range m {
		if v == p {
			continue
		}
		ret = append(ret, v)
		n--
		if n <= 0 {
			ret = append(ret, m[i+1:]...)
			break
		}
	}
	return ret
}

func getmax(m []int, pre int) int {
	for _, k := range m {
		if k >= pre {
			continue
		}
		return k
	}
	return 0
}

func s10(f int, l uint) int {
	for l-1 > 0 {
		f = f * 10
		l--
	}
	return f
}
