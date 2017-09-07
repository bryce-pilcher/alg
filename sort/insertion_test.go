package sort

import (
	"errors"
	"fmt"
	"github.com/bryce-pilcher/sliceutil"
	"testing"
)

func TestInsertion(t *testing.T) {
	cases := []struct {
		in, want []interface{}
	}{
		{[]interface{}{1, 3, 2}, []interface{}{1, 2, 3}},
		{[]interface{}{1, 3, 2, 10, 6, 4, 8, 7, 9, 5}, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		in := sliceutil.DeepCopy(c.in)
		Insertion(c.in, compare)
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
			t.Errorf("Insertion(%d)==%d, want %d\n", in, c.in, c.want)
		} else {
			fmt.Printf("Insertion(%d) == %d\n", in, c.want)
		}
	}
}

func compare(x interface{}, y interface{}) (int, error) {
	a := x.(int)
	b := y.(int)
	switch {
	case a < b:
		return 1, nil
	case a == b:
		return 0, nil
	case a > b:
		return -1, nil
	}
	return 0, errors.New("something went wrong in comparison")
}
