package main

import "fmt"

// can use ownerInfo: owner to embed the owner struct OR set the struct as a field
// to access attributes directly
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

// and this is how methods work in Go - param and value is how to read input
func (e gasEngine) getMilesLeft() uint8 {
	return e.mpg * e.gallons
}

// can also generalise functionality using interfaces, using another engine class and generic canMakeIt func
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) getMilesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	getMilesLeft() uint8
}

// and can pass this interface to the can make it function
// this lets it use an identical function signature for different types
// to do the same calculation
func canMakeIt(e engine, miles uint8) {
	if miles <= e.getMilesLeft() {
		fmt.Println("We can make it!")
	} else {
		fmt.Println("We can't make it!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{30, 10, owner{"John Doe"}}

	// anonymous struct
	var myOtherEngine = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	myEngine.mpg = 40
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Println(myOtherEngine.mpg, myOtherEngine.gallons)
	fmt.Printf("Miles left: %v\n", myEngine.getMilesLeft())
	fmt.Println("Checking interface usage")
	canMakeIt(myEngine, 75)
}
