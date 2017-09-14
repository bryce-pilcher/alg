package sort

import (
	"github.com/bryce-pilcher/util/slice"
)

func HeapSort(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	buildHeap(l, cmp)

	for i := len(l) - 1; i >= 1; i-- {
		slice.Swap(l, 0, i)
		heapify(l, 0, i, cmp)
	}
}

func buildHeap(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	n := len(l)

	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(l, i, n, cmp)
	}
}

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
