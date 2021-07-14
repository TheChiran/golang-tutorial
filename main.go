package main

import "fmt"

func main() {
	// range
	numbers := []int{1, 2, 3, 4, 5, 6}
	// _ is used to ignore index
	for index, number := range numbers {
		fmt.Print(index, ") ", number, "\n")
	}

}
