package main

import "fmt"

func recieveOnly(c <-chan string) {
	fmt.Println(<-c)
}
func sendOnly(c chan<- string) {
	c <- "Send only channel"
}
func main() {
	c := make(chan string, 1)
	c <- "Recieve Only Channel"
	recieveOnly(c)
	sendOnly(c)
	fmt.Println(<-c)
}
