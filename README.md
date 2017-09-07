# Algorithms in Go

This repository holds algorithms written in go for sorting, searching, and graph
manipulation.  

Variations may be included for performance benefits.  

The functions are written with interfaces such that any value can be passed into
the algorithm to be worked on. 

## Sort

These algorithms work on a slice of interfaces (`[]interface{}`) and mutate the
slice passed in. The caller is responsible for implementing and passing a 
compare function to the sort algorithm. This function will have to convert the
values to be compared from an interface to the proper Type.

Ex:
```go
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
	return 0, errors.New("Something went wrong in the comparison!")
}
```

### Insertion
This sorts an array like one might sort a hand of cards.  Start at the left side
of the hand, find the first item out of order, and insert that item where it 
belongs.

Runtime:  
Best - O(n)  
Average - O(n^2)  
Worst - O(n^2)  
