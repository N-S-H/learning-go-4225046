package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	//dayNumber := int(weekday)
	//fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	dayNumber := 0; switch dayNumber {
	case 1: 
	   result = "Its a monday"
	case 2:
		result = "Its a Tuesday"
	case 3:
		result = "Its a Wednesday"
	default:
		result = "Its a good day"			 
	}
	fmt.Println(result)
 
	x := 0
	switch {
	case x < 0:
		result = "Less than zero"
		fallthrough
	case x == 0:
		result = "Equals zero"
		fallthrough
	default:
		result = "Greater than zero" 
	}

}
