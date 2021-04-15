package main

import (
	"fmt"
	"math/rand"
)

/*

Save some money to buy a gift for your friend. Write a program that randomly places
nickels ($0.05), dimes ($0.10), and quarters ($0.25) into an empty piggy bank until it con-
tains at least $20.00. Display the running balance of the piggy bank after each deposit,
formatting it with an appropriate width and precision.

*/

func main() {
	const dime = 0.10
	const nikels = 0.05
	const quarters = 0.25
	piggy := 0.0

	for piggy <= 20.00 {

		choice := rand.Intn(3)

		switch choice {
		case 0:
			piggy += nikels
		case 1:
			piggy += dime
		case 2:
			piggy += quarters
		}

		fmt.Printf("$%5.2f\n", piggy)
	}
}
