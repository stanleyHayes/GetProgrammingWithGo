package main

import (
	"fmt"
	"math/rand"
)

/*

Write a new piggy bank program that uses integers to track the number of cents rather
than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢) into an empty
piggy bank until it contains at least $20.
Display the running balance of the piggy bank after each deposit in dollars (for exam-
ple, $1.05).

*/
func main() {
	const nikels = 5
	const dimes = 10
	const quarters = 25
	piggy := 0
	for piggy <= 2000 {
		choice := rand.Intn(3)
		switch choice {
		case 0:
			piggy += nikels
		case 1:
			piggy += dimes
		case 2:
			piggy += quarters
		}
		dollars := piggy / 100
		cents := piggy % 100

		fmt.Printf("$%02d.%02d\n", dollars, cents)
	}
}
