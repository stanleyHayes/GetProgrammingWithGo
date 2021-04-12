package main

import "fmt"

/*
Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. Write a program
to determine how fast a ship would need to travel (in km/h) in order to reach Malacan-
dra in 28 days. Assume a distance of 56,000,000 km.
*/
func main() {
	const hourPerDay = 24
	const days = 28
	var hours = hourPerDay * days
	var distance = 56000000
	var speed =  distance / hours
	fmt.Printf("The speed required for The Space Trilogy to get to Malacandra in 28 days is %v km/h", speed)
}
