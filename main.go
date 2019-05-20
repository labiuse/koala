package main

import (
	"container/heap"
	"fmt"
	"github.com/labiuse/koala/priorityqueue"
)

func main() {
	fmt.Println(FindRepeated([]int{1, 3, 3, 5, 5, 6, 6, 5, 3, 3}, 2))
}

// FindRepeated returns n top most occured number in the list of s numbers
func FindRepeated(s []int, n int) []int {

	var (
		// use map to keep track of occurunces in o(1) instead of iterating over
		// slice of pq
		indexes = make(map[int]int)
		result  = []int{}
		pq      priorityqueue.PriorityQueue
	)

	// Early exit to prevent later panic caused by calling heap.Pop if pq is empty
	if len(s) <= 1 {
		return result
	}

	for _, v := range s {
		if index, ok := indexes[v]; !ok {
			indexes[v] = len(pq)
			pq = append(pq, &priorityqueue.Item{
				Value:    v,
				Priority: 1,
				Index:    indexes[v],
			})
		} else {
			pq[index].Priority++
		}
	}

	heap.Init(&pq)

	for i := 0; i < n; i++ {
		if item, ok := heap.Pop(&pq).(*priorityqueue.Item); ok {
			result = append(result, item.Value)
		}
	}

	return result
}
