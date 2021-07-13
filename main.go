package main

import "fmt"

func main() {
	//go-lang variables
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}
