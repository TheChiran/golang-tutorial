package main

import "fmt"

func main() {
	// define array
	a := [5]int{1, 2, 3, 4, 5}

	// create slice from array
	t := a[1:3:5]

	// output slice
	fmt.Println(t)
}
