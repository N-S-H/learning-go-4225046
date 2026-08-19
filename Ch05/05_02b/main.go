package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	fmt.Printf("The %v says %v!\n", dog.Breed, dog.Sound)
	dog.Speak()
}

type Dog struct {
   Breed string
   Sound string
}

func (d Dog) Speak() {
   fmt.Println("The dog",d.Breed,"says sound: ",d.Sound)
}
