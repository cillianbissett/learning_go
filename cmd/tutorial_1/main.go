package main

import "fmt"

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)
	var a, b int = 10, 3
	var result, remainder, err = intDivision(a, b)
	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Println("Result of division is", result)
	case remainder != 0:
		fmt.Println("Result of division is", result, "with remainder", remainder)
	}
	switch remainder {
	case 0:
		fmt.Println("Division was exact")
	case 1, 2:
		fmt.Println("Division was not exact")
	default:
		fmt.Println("Division was miles off")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = fmt.Errorf("Division by zero")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
