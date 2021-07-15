package main

import (
	"fmt"
)

type House struct {
	rooms int
	price int32
	city  string
}

func main() {
	// struct
	var house House
	house.rooms = 4
	house.price = 36550
	house.city = "Dhaka"

	fmt.Printf("House.rooms = %d \n", house.rooms)
	fmt.Printf("House.price = %d \n", house.price)
	fmt.Printf("House.city = %s \n", house.city)

}
