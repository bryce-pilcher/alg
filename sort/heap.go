package sort

import (
	"github.com/bryce-pilcher/util/slice"
)

/*
   This function is the entry point to the heap sort algorithm.

   It kicks things off by building the heap, then recurses the problem,
   operating on smaller subarrays each time.
*/
func HeapSort(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	buildHeap(l, cmp)

	for i := len(l) - 1; i >= 1; i-- {
		slice.Swap(l, 0, i)
		heapify(l, 0, i, cmp)
	}
}

/*
   This function takes the original slice and processes it so it resembles a heap.
*/
func buildHeap(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	n := len(l)

	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(l, i, n, cmp)
	}
}

/*
   heapify takes a slice and recursively reorders it to resemble a heap.

   It operates on a slice from 0 - max, which allows it to operate on smaller
   and smaller pieces of the heap.
*/
func heapify(l []interface{}, index int, max int, cmp func(x interface{}, y interface{}) (int, error)) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < max {
		c, err := cmp(l[left], l[index])
		checkerror(err, false)
		if c > 0 {
			largest = left
		}
	}

	if right < max {
		c, err := cmp(l[right], l[largest])
		checkerror(err, false)
		if c > 0 {
			largest = right
		}
	}

	if largest != index {
		slice.Swap(l, index, largest)
		heapify(l, largest, max, cmp)
	}
}
