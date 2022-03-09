package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

// var testSlice = []int{3, 5, 5, 9, 11}
// var testSlice1 = []int{1, 3, 5, 7}

func main() {
	path := ""
	fmt.Print("Provide path to a file of integers, please: ")

	// taking user input
	fmt.Scanln(&path)
	startTime := time.Now()
	fmt.Println("Time Now: ", startTime)
	time.Sleep(2 * time.Second)

	// Open file and create scanner on top of it
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanLines)

	i := 0
	min := 0
	max := 0
	sum := 0
	arithmeticMean := 0
	sortedSlice := make([]int, 0)

	// Scan for next token
	for scanner.Scan() {
		if scanner.Text() != "" {
			if n, err := strconv.Atoi(scanner.Text()); err == nil {
				if n < min {
					min = n
				}
				if n > max {
					max = n
				}
				sum += n
				i++
				sortedSlice = append(sortedSlice, n)

			}

		}
		continue
	}
	sort.Ints(sortedSlice)

	median := Median(i, sortedSlice)
	arithmeticMean = sum / i

	fmt.Printf("Min: %d\nMax: %d\n", min, max)
	fmt.Printf("Sum: %d \n", sum)
	fmt.Printf("Arithmetic mean: %d \n", arithmeticMean)
	fmt.Printf("Median: %d \n", median)

	// False on error or EOF. Check error
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Scan completed and reached EOF.")
	}

	fmt.Println("Time Since: ", time.Since(startTime))

}

func Median(elemNum int, sortedSlice []int) int {
	median := 0
	if elemNum%2 != 0 {
		median = sortedSlice[elemNum/2]
	} else {
		median = (sortedSlice[elemNum/2-1] + sortedSlice[elemNum/2]) / 2
	}
	return median
}
