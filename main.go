package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var fileName string

	fmt.Println("What is the file name?")
	_, err := fmt.Scan(&fileName)
	check(err)
	//fmt.Printf("%s", fileName)

	dat, err := os.Open(fileName)
	check(err)

	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		fmt.Printf("line: %s \n", scanner.Text())
	}

}

func countWords() {

}

func reportResults() {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
