package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string

	fmt.Println("What is the file name?")
	_, err := fmt.Scan(&fileName)
	check(err)
	//fmt.Printf("%s", fileName)

	dat, err := os.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))

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
