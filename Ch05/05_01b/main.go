package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something")
	value1 := 5
	value2 := 10
	value3 := 56
	sum := addValues(value1,value2)
	fmt.Printf("The sum of %v and %v is %v\n",value1,value2,sum)

	sum = addAllValues(value1,value2,value3)
	fmt.Println("The sum is :",sum)

	const der int = 35
	fmt.Printf("The number is %v",der)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
    sum := 0
		for _, v := range values {
			sum += v
		}
		return sum
}
