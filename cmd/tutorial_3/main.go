package main

import "fmt"

func main() {
	var myString = "Sample"
	// var indexed = myString[0]
	for i, v := range myString {
		fmt.Printf("At index %v, value is %v\n", i, v)
	}
}
