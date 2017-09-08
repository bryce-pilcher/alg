package sort

import (
	"github.com/bryce-pilcher/util/rand"
)

/*
   This function performs the quicksort algorithm on a slice.

   Quicksort is like Median sort, but simpler.  It cuts
   out finding the exact median of the slice.  This can lead
   to uneven splits whereas median sort quarantees even splits.
*/
func Quick(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	right := len(l) - 1
	quicksort(l, 0, right, cmp)
}

func quicksort(l []interface{}, left int, right int, cmp func(x interface{}, y interface{}) (int, error)) {
	if right <= left {
		return
	}
	pivotIdx := rand.RangeInt(left, right)
	pivotIdx = partition(l, left, right, pivotIdx, cmp)

	quicksort(l, left, pivotIdx-1, cmp)
	quicksort(l, pivotIdx+1, right, cmp)
}
