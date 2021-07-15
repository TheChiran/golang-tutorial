package main

import (
	"fmt"
)

func main() {
	// golang map
	elements := make(map[string]string)
	elements["O"] = "Oxygen"
	elements["Ca"] = "Calcium"
	elements["C"] = "Carbon"

	fmt.Println(elements["C"])
}
