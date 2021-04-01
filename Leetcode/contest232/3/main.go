package main

import (
	"container/heap"
)

type maxHeap []Pair

type Pair struct {
	Index int
	C     float64
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].C > h[j].C
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(num interface{}) {
	*h = append(*h, num.(Pair))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	q := &maxHeap{}
	for i, v := range classes {
		a := float64(v[0]) / float64(v[1])
		b := float64(v[0]+1) / float64(v[1]+1)
		heap.Push(q, Pair{
			Index: i,
			C:     b - a,
		})
	}
	heap.Init(q)
	for i := 0; i < extraStudents; i++ {
		temp := heap.Pop(q).(Pair)
		classes[temp.Index][0]++
		classes[temp.Index][1]++
		a := float64(classes[temp.Index][0]) / float64(classes[temp.Index][1])
		b := float64(classes[temp.Index][0]+1) / float64(classes[temp.Index][1]+1)
		temp.C = b - a
		heap.Push(q, temp)
	}
	var sum float64
	for _, v := range classes {
		sum += float64(v[0]) / float64(v[1])
	}
	return sum / float64(len(classes))
}
