package main

import "fmt"

func arrayExample() {
	fmt.Println("Array example")
	var array = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)
}
func main() {
	//arrays
	//in golang array value is set to default 0
	var array = new([5]int)
	fmt.Println(array[2])
	fmt.Println(array[3])
	arrayExample()
	arrayString()
}
func arrayString() {
	var string = []string{"T", "O", "N", "M", "O", "Y"}
	fmt.Println("Array string example")
	fmt.Println(string)
}
