package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//go-lang keyboard input
	// inputReader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// name, _ := inputReader.ReadString('\n')
	// fmt.Print("Your name is: " + name)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number: ")
	number, _ := inputReader.ReadString('\n')

	number = strings.Replace(number, "\n", "", -1)

	num, e := strconv.Atoi(number)

	if e != nil {
		fmt.Println("Conversion error: ", number)
	}

	if num >= 1 && num <= 10 {
		fmt.Println("Correct")
	} else {
		fmt.Println("Num is not in range")
	}

}
