package main

import (
	"fmt"
	"os"
)

func main() {
	//flie exists
	// if _, err := os.Stat("go.txt"); err == nil {
	// 	fmt.Println("File exists\n")
	// } else {
	// 	fmt.Println("File does not exists\n")
	// }
	if _, err := os.Stat("file-exists.file"); os.IsNotExist(err) {
		fmt.Println("File does not exists")
	}
}
