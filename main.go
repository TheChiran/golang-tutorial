package main

import "fmt"

func values() (int, int) {
	return 2, 4
}

func main() {
	//returning multiple values
	x, y := values()
	fmt.Println(x)
	fmt.Println(y)
}
