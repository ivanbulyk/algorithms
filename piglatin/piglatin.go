package main

import (
	"fmt"
	"strings"
)

const (
	str string = "Do not communicate by sharing memory; instead, share memory by communicating."
)

func interpret(input string) string {

	var pigLatWords []string
	enWords := strings.Split(input, " ")

	for _, word := range enWords {
		firstLetter := word[0:1]
		if strings.Contains("aeiou", firstLetter) {
			pigLatWords = append(pigLatWords, word+"ay")
		} else if strings.Contains("AEIOU", firstLetter) {
			pigLatWords = append(pigLatWords, word+"ay")
		} else {
			pigLatWords = append(pigLatWords, word[1:]+firstLetter+"ay")
		}
	}
	return strings.Join(pigLatWords, " ")
}

func main() {

	fmt.Println(interpret(str))
}
