package main

import "fmt"

func makeSequence() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}

func main() {
	sequenceGenerator := makeSequence()
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
	fmt.Println(sequenceGenerator())
}
