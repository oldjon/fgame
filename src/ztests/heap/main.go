package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

//Heap Interface
//type Interface interface {
//    sort.Interface
//    Push(x interface{}) // 向末尾添加元素
//    Pop() interface{}   // 从末尾删除元素
//}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Sort Interface
//type Interface interface {
//    // Len方法返回集合中的元素个数
//    Len() int
//    // Less方法报告索引i的元素是否比索引j的元素小
//    Less(i, j int) bool
//   // Swap方法交换索引i和j的两个元素
//    Swap(i, j int)
//}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	h := &IntHeap{5, 8, 1, 3, 4}
	heap.Init(h)
	heap.Push(h, 6)

	fmt.Printf("minimum: %d\n", (*h)[0])
	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println("aaaaa")

	//output :
	//minimum: 1
	//1 3 4 5 6 8
}
