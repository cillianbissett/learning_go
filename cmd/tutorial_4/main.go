package main

import "fmt"

type gasEngine struct {
	// fields
	mpg     uint8
	gallons uint8
}

type owner struct {
}

func main() {
	// can omit names
	var myEngine gasEngine = gasEngine{mpg: 30, gallons: 10}

	// can overwrite
	myEngine.mpg = 40
	fmt.Println(myEngine.mpg, myEngine.gallons)
}
