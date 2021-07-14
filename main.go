package main

import "fmt"

//when a variable is assign insides a function
//that is local variable
//when that variable is assigned outside of function
//that is global
var x = 7 //local variables

func example() {
	fmt.Println(x)
}
func main() {
	//scopes
	fmt.Println(x)
	example()
}
