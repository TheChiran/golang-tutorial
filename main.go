package main

import "fmt"

//a recursion is something that calls self
//and have a condition to stop program
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	x := factorial(3)
	fmt.Println(x)
}
