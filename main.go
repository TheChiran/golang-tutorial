package main

import "fmt"

func main() {
	/* define a string */
	mystring := "Hello"

	/* take slice */
	string1 := mystring[0:2]
	string2 := mystring[2:5]
	fmt.Println(string1)
	fmt.Println(string2)
}
