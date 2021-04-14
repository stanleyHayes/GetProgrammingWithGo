package main

import (
	"fmt"
	"math/rand"
)

/*

Generate a random year instead of always using 2018 .
 For February, assign daysInMonth to 29 for leap years and 28 for other years.
Hint: you can put an if statement inside of a case block.
 Use a for loop to generate and display 10 random dates.
*/
func main() {
	era := "AD"
	daysInMonth := 31

	for i := 0; i < 10; i++ {
		month := rand.Intn(12) + 1
		year := rand.Intn(9999)

		switch month {
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				daysInMonth = 29
			}
			daysInMonth = 28
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, day, month, year)
	}
}
