package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func who() {
	fmt.Println("Go")
}

func main() {
	defer who()
	hello()
}
