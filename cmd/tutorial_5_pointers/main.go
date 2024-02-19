package main

import "fmt"

// use case for halving memory footprint
// dont pass arrays to funcs, pass the pointers underpinning it
func square(thing2 *[5]int32) [5]int32 {
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	fmt.Printf("Memory address of array in func: %p\n", thing2)
	return *thing2
}

func main() {
	var p *int32 = new(int32)
	var i int32 = 42

	// note use of * in front of p to get the value!
	fmt.Printf(("The value p points to is: %v\n"), *p)
	fmt.Printf(("The value of i is: %v\n"), i)

	// really spicy: say we have a var for a pointer, which points to an address of another var
	// if we update the value of the pointer, the value of the original var will also change
	var a int32 = 5
	var b *int32 = &a
	*b = 10
	fmt.Printf("The value of a is: %v\n", a)

	var thing1 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("Original array memory address: %p", &thing1)
	var result [5]int32 = square(&thing1)
	fmt.Printf("Squared array: %v", result)

}
