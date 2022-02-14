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

	slices := strings.Split(string(processedString), " ")

	//slices := regexp.MustCompile("[-,' '\n]").Split(string(dat), -1)

	//for _, slice := range slices {
	//	fmt.Print(slice)
	//}

	wordCountMap, err := countWords(slices)
	check(err)
	reportResults(wordCountMap)
}

func countWords(slices []string) (map[string]int, error) {
	wordCountMap := make(map[string]int)

	for _, word := range slices {
		//word, err := removePunc(word)
		//if err != nil {
		//	return nil, err
		//}
		if _, ok := wordCountMap[word]; ok {
			wordCountMap[word] += 1
		} else {
			wordCountMap[word] = 1
		}
	}
	return wordCountMap, nil
}

func reportResults(data map[string]int) {
	for k, v := range data {
		fmt.Println(k, "value is ", v)
	}
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

	processedString := reg.ReplaceAllString(unprocessedString, " ")

	return processedString, err
}
