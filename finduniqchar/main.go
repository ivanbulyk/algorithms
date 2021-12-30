package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	strFinal  string
	charFinal string
)

func userInput() []string {
	outputSlice := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your text please. For finishing your input just hit 'Enter', without any value: ")
	for {

		scanner.Scan()

		text := scanner.Text()

		if len(text) != 0 {

			outputSlice = append(outputSlice, strings.Split(text, " ")...)
		} else {
			break
		}
	}

	return outputSlice
}

func findUniqChar() string {
	inpuWords := userInput()
	for _, word := range inpuWords {
		for i := 0; i < len(word); i++ {
			if strings.Contains(word[i+1:], string(word[i])) || unicode.IsPunct(rune(word[i])) {
				continue
			} else {
				strFinal = strFinal + string(word[i])
				break
			}
		}
	}

	for i := 0; i < len(strFinal); i++ {
		if strings.Contains(strFinal[i+1:], string(strFinal[i])) {
			continue
		} else {
			charFinal = charFinal + string(strFinal[i])
			break
		}
	}
	return charFinal
}

func main() {

	fmt.Println("Your result is: ", findUniqChar())
}
