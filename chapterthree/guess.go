package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var secret = rand.Intn(100) + 1
	var count = 0
	for {
		var guess = 0
		fmt.Printf("Guess a number between 1 and 100 inclusive: ")
		_, err := fmt.Scanf("%d", &guess)
		count++
		if err != nil {
			panic(err)
		}
		if guess > secret {
			fmt.Printf("%d is bigger than secret number\n", guess)
		} else if guess < secret {
			fmt.Printf("%d is smaller than secret number\n", guess)
		} else if secret == guess {
			fmt.Printf("You guessed the secret number, %d right at the guess number %d\n", secret, count)
			break
		}
	}
}
