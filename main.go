// Connor Henderson

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName string

	fmt.Println("What is the file name?")
	_, err := fmt.Scan(&fileName)
	check(err)

	dat, err := os.ReadFile(fileName)
	check(err)

	processedString, err := removePunc(string(dat))
	check(err)

	slices := strings.Split(processedString, " ")

	//for _, slice := range slices {
	//	fmt.Print(slice)
	//}

	countWords(slices)
}

func countWords(slices []string) map[string]int {
	wordCountMap := make(map[string]int)

	for _, word := range slices {
		if _, ok := wordCountMap[word]; ok {
			wordCountMap[word] += 1
		} else {
			wordCountMap[word] = 1
		}
	}
	return wordCountMap
}

func reportResults() {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func removePunc(unprocessedString string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}

	processedString := reg.ReplaceAllString(unprocessedString, "")

	return processedString, err
}
