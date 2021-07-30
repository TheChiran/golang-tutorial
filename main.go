package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func main() {
	users := []User{
		{"Debora", "12345"},
		{"Bob", "1345"},
		{"Sandra", "12"},
	}

	out, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
