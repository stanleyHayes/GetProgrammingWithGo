package main

import "fmt"

/*
Write a program that converts strings to Booleans:
 The strings “true”, “yes”, or “1” are true .
 The strings “false”, “no”, or “0” are false .
 Display an error message for any other values.
TIP The switch statement accepts multiple values per case, as covered in lesson 3.
*/
func main() {
	var response string
	var boolean bool
	fmt.Printf("Enter a value to be converted to boolean:\t")
	_, err := fmt.Scanf("%s", &response)
	if err != nil {
		panic(err)
	}
	switch response {
	case "true", "yes", "1":
		boolean = true
	case "false", "no", "0":
		boolean = false
	default:
		panic("Unknown value")
	}
	fmt.Printf("Response %v is %v", response, boolean)
}
