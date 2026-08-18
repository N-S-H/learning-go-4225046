package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	fmt.Println(str1,str2,str3)

	aNumber := 42
	strLen, err := fmt.Println("The value is",aNumber)
	if err == nil {
		fmt.Println("The length is",strLen)
	}

	fmt.Printf("Value of the number %v\n",aNumber)
	fmt.Printf("Datatype of value is %T\n",aNumber)
}
