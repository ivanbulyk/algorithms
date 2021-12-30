package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func Counter(s string) string {
	// Write your code here

	strSorted := SortStringByCharacter(s)
	r := StringToRuneSlice(strSorted)
	m := make(map[rune]int)

	for _, v := range r {
		value, ok := m[v]
		if ok {
			value++
			m[v] = value
		} else {
			m[v] = 1
		}
	}

	str := ""
	for i := 0; i < len(r); i++ {
		for k, v := range m {
			if r[i] == k {
				str += fmt.Sprint(string(k)) + fmt.Sprint(v)
				delete(m, r[i])
			}
		}

	}

	return str
}

func main() {
	// strInput := strings.Join(os.Args[1:], "")
	// strSorted := SortStringByCharacter(strInput)
	strCounted := Counter(strings.Join(os.Args[1:], ""))

	fmt.Println(strCounted)

}
