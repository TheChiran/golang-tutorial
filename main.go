package main

import "fmt"

func main() {
	i := 1
	max := 20

	//technically go dosent have while
	//for can be used while in go

	for i < max {
		fmt.Println(i)
		i++
	}
}
