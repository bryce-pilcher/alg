package sort

import (
	"fmt"
	"github.com/bryce-pilcher/util/slice"
)

/*
   partition separates the array into two halves around a pivot point.
   The values before the pivot point in the slice are also less than the
   value at the pivot point, although not necessarily in order.  Same for
   the values after the pivot point in the slice, the values are greater.
*/
func partition(l []interface{}, left int, right int, idx int, cmp func(x interface{}, y interface{}) (int, error)) (store int) {
	store = left
	piv := l[idx]
	slice.Swap(l, idx, right)

	for i := left; i < right; i++ {
		c, err := cmp(l[i], piv)
		if err != nil {
			fmt.Println(err)
			break
		}

		if c < 0 {
			slice.Swap(l, store, i)
			store = store + 1
		}
	}

	slice.Swap(l, store, right)
	return
}

/*
   Checks an error and panics if the err is not nil and p is true, otherwise
   just prints the err if not nil.
*/
func checkerror(err error, p bool) {
	if err != nil {
		if p {
			panic(err)
		} else {
			fmt.Println(err)
		}
	}
}
