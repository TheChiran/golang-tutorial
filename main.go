package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	// golang random numbers
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 6)
	fmt.Println("Random Number: ", randomNum)

}
