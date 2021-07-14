package main

import "fmt"

func main() {
	// for loops
	for x := 0; x < 4; x++ {
		fmt.Printf("Iteration x: %d\n", x)
	}
	loopArray()
	countFromOneToTen()
}
func loopArray() {
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Loop Array")
	for i := 0; i < len(array); i++ {
		fmt.Println("Number : ", array[i])
	}
}

func countFromOneToTen() {
	fmt.Println("Count From one to ten")
	for i := 0; i < 10; i++ {
		fmt.Println("Count : ", (i + 1))
	}
}
