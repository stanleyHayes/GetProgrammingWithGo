package main

import "fmt"

func main() {
	result := kelvinToFahrenheit(0)
	fmt.Printf("%.2f K in F is %.2f", 0.0, result)
}

func celsiusToFahrenheit(c float64) float64{
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToCelsius(k float64) float64{
	return k - 273.15
}

func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	return celsiusToFahrenheit(c)
}