package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var start = 10
	for start > 0 {
		fmt.Println(start)
		if rand.Intn(100) == 0 {
			break
		}
		start--
	}
	if start == 0 {
		fmt.Println("Lift Off!!")
	}else {
		fmt.Println("Launch Failed")
	}
}
