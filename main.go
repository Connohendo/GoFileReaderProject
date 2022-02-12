package main

import (
	"fmt"
)

func main() {
	var fileName string

	fmt.Println("What is the file name?")
	_, err := fmt.Scan(&fileName)
	check(err)
	//fmt.Printf("%s", fileName)

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
