package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an slice
	var colors = []string{"Red", "Green", "Blue"}

	var colorsMake = make([]string, 0, 3)
  colorsMake = append(colorsMake, "Red", "Blue","Violet","Purple")
	fmt.Println(colorsMake) 
	fmt.Println(colors)

	colorsMake = remove(colorsMake, 1)
	fmt.Println(colorsMake)

	sort.Strings(colorsMake)
	fmt.Println(colorsMake)
}

func remove(slice []string, i int) []string {
  return append(slice[:i], slice[i+1:]...)
} 
