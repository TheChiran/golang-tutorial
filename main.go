package main

import (
	"fmt"
)

func main() {
	// golang store information on map
	website := map[string]map[string]string{
		"Google": map[string]string{
			"name": "Google",
			"type": "Search",
			"work": "Search Engine",
		},
		"YouTube": map[string]string{
			"name": "YouTube",
			"type": "Video",
			"work": "Video Search Engine",
		},
	}

	fmt.Println(website["Google"]["name"])
	fmt.Println(website["YouTube"]["work"])
}
