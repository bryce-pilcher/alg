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
		return -1, nil
	case a == b:
		return 0, nil
	case a > b:
		return 1, nil
	}
	return 0, errors.New("Something went wrong in the comparison!")
}
```

### Insertion
This sorts a slice like one might sort a hand of cards.  Start at the left side
of the hand, find the first item out of order, and insert that item where it 
belongs.

Runtime:  
Best - O(n)  
Average - O(n^2)  
Worst - O(n^2)  

### Median
Median sort takes a slice and looks through it to find the median value.  As it
is doing this, it is also partitioning the slice around a pivot index.  After
a successful pivot, the values before the pivot are less than the pivot, though
not necessarily in order.  The same is true of the values after the pivot.

Once it finds the median, it recurses the problem into  smaller subsets and 
repeats the process.

Runtime:  
Best - O(n log n)  
Average - O(n log n)  
Worst - O(n^2)

### Quicksort
Quicksort is similar to Median sort, but it does not care about the median
value, so it skips finding the exact median value and just recurses the smaller
problems.  There are many variations which can enable a quicker sorting time.

Variations will be built in eventually. These include using a stack rather than 
recursion for tracking sub problems.  

Runtime:  
Best - O(n log n)  
Average - O(n log n)  
Worst - O(n^2)

### Heap Sort
Heap sort works differently than all of the above algorithms.  Instead of 
operating directly on the slice, it converts the slice to a heap and then sorts
the heap. You may notice, there is no heap data structure used; it reorders the
slice in a way that represents a heap.  This aids the function in simply 
copying the largest value to the end of the array, and then re-heaping the new
subarray.  

Runtime:  
Best - O(n log n)  
Average - O(n log n)  
Worst - O(n log n)
