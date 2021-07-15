package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//read files
	b, err := ioutil.ReadFile("go.txt")

	if err != nil {
		fmt.Println(err)
	}

	string := string(b)

	fmt.Println(string)

}
