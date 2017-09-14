package sort

import (
	"fmt"
	"github.com/bryce-pilcher/util/slice"
	"testing"
)

// compare used in this function is found in insertion_test.go
func TestHeap(t *testing.T) {
	cases := []struct {
		in, want []interface{}
	}{
		{[]interface{}{1, 3, 2}, []interface{}{1, 2, 3}},
		{[]interface{}{1, 3, 2, 10, 6, 4, 8, 7, 9, 5}, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		in := slice.DeepCopy(c.in)
		HeapSort(c.in, compare)
		pass := false
		for i, n := range c.in {
			if n == c.want[i] {
				pass = true
			} else {
				pass = false
				break
			}
		}
		if !pass {
			t.Errorf("HeapSort(%d)==%d, want %d\n", in, c.in, c.want)
		} else {
			fmt.Printf("HeapSort(%d) == %d\n", in, c.want)
		}
	}
}
