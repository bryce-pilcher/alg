package sort

import (
	"github.com/bryce-pilcher/util/rand"
)

/*
   This function is designed to perform a median sort on a slice of any data type
   The caller is responsible for passing a compare function that handles the
   interface value.

   Median sort works by finding a value the median value in an array and
   partitioning that array into smaller and smaller subsets to solve the
   problem.
*/
func Median(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	right := len(l) - 1
	mediansort(l, 0, right, cmp)
}

/*
   A recursive function that handles the partitioning of the array into smaller sizes
   The return from selectKth is not actively used, that is why it is ignored.
*/
func mediansort(l []interface{}, left int, right int, cmp func(x interface{}, y interface{}) (int, error)) {
	if left < right {
		mid := (right - left + 1) / 2
		_ = selectKth(l, mid+1, left, right, cmp)
		mediansort(l, left, left+mid-1, cmp)
		mediansort(l, left+mid+1, right, cmp)
	}
}

/*
   selectKth finds the index of the median element in the slice.
   This works with partition to find the median value in a slice.
*/
func selectKth(l []interface{}, mid int, left int, right int, cmp func(x interface{}, y interface{}) (int, error)) (pivotIdx int) {
	// Go does not provide a function for getting a random value
	// from a range, so this bounds the random value between
	// left and right
	idx := rand.RangeInt(left, right)
	pivotIdx = partition(l, left, right, idx, cmp)
	if (left + mid - 1) == pivotIdx {
		return
	}

	if (left + mid - 1) < pivotIdx {
		return selectKth(l, mid, left, pivotIdx-1, cmp)
	} else {
		return selectKth(l, mid-(pivotIdx-left+1), pivotIdx+1, right, cmp)
	}
}
