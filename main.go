package main

import "fmt"

func main() {
	c := make(chan int, 5)

	c <- 5
	c <- 3

	fmt.Println(<-c)
	close(c)
	fmt.Println(<-c)
}
