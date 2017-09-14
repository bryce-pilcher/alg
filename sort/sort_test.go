package sort

import (
	"errors"
	"fmt"
	"github.com/bryce-pilcher/util/slice"
	"reflect"
	"runtime"
	"testing"
)

func TestInsertion(t *testing.T) {
	cases := []struct {
		in, want []interface{}
	}{
		{[]interface{}{1, 3, 2}, []interface{}{1, 2, 3}},
		{[]interface{}{1, 3, 2, 10, 6, 4, 8, 7, 9, 5}, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	functions := []interface{}{Insertion, Median, Quick, HeapSort}

	for _, c := range cases {
		in := slice.DeepCopy(c.in)
		for _, f := range functions {
			f.(func(l []interface{}, cmp func(x interface{}, y interface{}) (int, error)))(c.in, compare)
			pass := false
			for i, n := range c.in {
				if n == c.want[i] {
					pass = true
				} else {
					pass = false
					break
				}
			}
			name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			if !pass {
				t.Errorf("%s(%d)==%d, want %d\n", name, in, c.in, c.want)
			} else {
				fmt.Printf("%s(%d) == %d\n", name, in, c.want)
			}
		}
	}
}

func compare(x interface{}, y interface{}) (int, error) {
	a := x.(int)
	b := y.(int)
	switch {
	case a < b:
		return -1, nil
	case a == b:
		return 0, nil
	case a > b:
		return 1, nil
	}
	return 0, errors.New("something went wrong in comparison")
}
