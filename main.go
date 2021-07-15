package main

import (
	"os"
)

func main() {
	// write files
	file, err := os.Create("file.txt")

	if err != nil {
		return
	}
	var string = "Write file in golang"

	defer file.Close()
	file.WriteString(string)
}
