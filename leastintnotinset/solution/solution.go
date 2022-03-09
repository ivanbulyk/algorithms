package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	max := 0
	sort.Ints(A)
	for _, val := range A {
		if val > max {
			if val-max > 1 && max > 0 {
				return max + 1
			}
			max = val
		}
	}
	return max + 1
}
