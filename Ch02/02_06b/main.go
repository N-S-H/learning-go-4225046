package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n",t)

	now := time.Now()
	fmt.Printf("The time currently is %v\n",now)

	fmt.Printf(now.Format(time.ANSIC)+"\n")

	format := "Mon 2006-02-01"
	fmt.Println(now.AddDate(0,0,1).Format(format)+"\n")

}
