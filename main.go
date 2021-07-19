package main

import "fmt"

func f(c chan string) {
	c <- "f() was here"
}

func f2(c chan string) {
	msg := <-c
	fmt.Println("f2", msg)
}

func main() {
	//go routiens
	//channels are a way to communicate with each other
	var c chan string = make(chan string)
	go f(c)
	go f2(c)

	fmt.Scanln()
}
