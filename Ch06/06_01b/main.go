package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	name := "./fromString.txt"
	file, err := os.Create(name)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file,"Hello from Go!")
	fmt.Printf("Wrote a file with character count: %v\n",length)
  readFile(name)
}

func readFile(name string) {
	data, err := os.ReadFile(name)
	checkError(err)
	fmt.Println("Text read from file",string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}