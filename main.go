package main

import (
	"errors"
	"fmt"
)

func do() (int, error) {
	return -1, errors.New("Something wrong")
}

func main() {
	fmt.Println(do())
}
