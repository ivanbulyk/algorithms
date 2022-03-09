package main

import (
	"fmt"
	"sort"

	"github.com/ivanbulyk/algorithms/leastintnotinset/solution"
)

// var testSlice = []int{3, 1, 5, 9, 11}
var testSlice = []int{1, 0, -1}

func main() {

	sort.Ints(testSlice)
	fmt.Println(testSlice)
	sol := solution.Solution(testSlice)
	fmt.Println(sol)

}
