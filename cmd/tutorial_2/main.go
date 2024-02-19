package main

import "fmt"

func main() {
	// arrays are fixed length
	// in contiguous memory
	intArr := [5]int32{1, 2, 3, 4, 5}
	fmt.Println(intArr)
	intArr[1] = 10
	fmt.Println(intArr[0:4])

	// slices are wrappers, not fixed in length
	// not neccesarily contiguous if we append
	intSlice := []int32{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 6)
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{7, 8, 9}
	bigSlice := append(intSlice, intSlice2...)
	fmt.Println(bigSlice)

	// can specify up front capacity to avoid reallocation of memory
	// reallocation is expensive
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// maps are key value pairs (Python dictionaries)
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"Bagpipes": 1, "Drums": 2, "Guitar": 3}
	fmt.Println(myMap2["Bagpipes"])

	// returns optional boolean to identify if value is there
	// can delete by reference / key
	// KEYS NOT ORDERED FOR ITERATION!

	for name := range myMap2 {
		fmt.Println(name, myMap2[name])
	}

	for i, v := range intSlice3 {
		fmt.Printf("At index %v, value is %v\n", i, v)
	}

	// no while loops per se
	var i int = 0
	for i < 5 {
		fmt.Println("Iteration", i)
		i++
	}

	// fancier implementation of while loop
	for j := 0; j < 5; j++ {
		fmt.Println("Iteration", j)
	}
}
