package main

import (
	"fmt"
	"math/rand"
)

/*

The table should have four columns:
 The spaceline company providing the service
 The duration in days for the trip to Mars (one-way)
 Whether the price covers a return trip
 The price in millions of dollars 
For each ticket, randomly select one of the following spacelines: Space Adventures,
SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km
away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine
the duration for the trip to Mars and also the ticket price. Make faster ships more expen-
sive, ranging in price from $36 million to $50 million. Double the price for round trips.

*/
func main() {
	const distance = 62100000
	const hoursPerDay = 24
	const secondsPerHour = 3600
	fmt.Printf("%-25v%-10v%-20v%-5v\n", "Space Line", "Days", "Trip Type", "Price")
	fmt.Printf("====================================================================\n")
	for i := 0; i < 10; i++ {
		spaceLineNumber := rand.Intn(3)
		var spaceLine = "Virgin Galactic"
		switch spaceLineNumber {
		case 0:
			spaceLine = "Space Adventures"
		case 1:
			spaceLine = "SpaceX"
		case 2:
			spaceLine = "Virgin Galactic"
		}

		speedInKmPerS := rand.Intn(15) + 16
		durationInSeconds := distance/speedInKmPerS
		durationInDays := durationInSeconds / (hoursPerDay * secondsPerHour)
		price := speedInKmPerS + 20
		tripType := "One-Way"
		trip := rand.Intn(2)
		if trip == 1 {
			price *= 2
			tripType = "Round-Trip"
		}

		fmt.Printf("%-25v%-10v%-20v $%+5v\n", spaceLine, durationInDays, tripType, price)
	}

}
