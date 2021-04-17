package main

import "fmt"

func main() {
	var word = "shalom"
	for _, character := range word {
		fmt.Printf("%c\n", character)
	}
}
