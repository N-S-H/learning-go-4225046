package main

import (
	"fmt"
	"strings"
)

func main() {
	len,_ := fmt.Println("Conditional logic")
  if strings.Contains(string(len),"1") {
		fmt.Println("Length probably in range 10-19")
	} else {
		fmt.Println("Our guess is wrong")
	}
}
