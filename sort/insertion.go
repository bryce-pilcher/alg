package sort

import (
	"fmt"
)

/*
   Function for performing an insertion sort.
   Requires a comparator function to be passed into it.
   Insertion sort works well for small slices of data.
   It's runtime is:
	Best: O(n)
     Average: O(n^2)
       Worst: O(n^2)
*/
func Insertion(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	n := len(l)
	for i := 1; i < n; i++ {
		insert(l, i, l[i], cmp)
	}
}

// This function performs the insertion of the object at the correct location
func insert(l []interface{}, pos int, value interface{}, cmp func(x interface{}, y interface{}) (int, error)) {
	i := pos - 1
	for i >= 0 {
		c, err := cmp(value, l[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		if c < 0 {
			l[i+1] = l[i]
			i = i - 1
		} else {
			break
		}
	}
	l[i+1] = value
}
