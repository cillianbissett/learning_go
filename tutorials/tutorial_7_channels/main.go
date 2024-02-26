package main

import "fmt"

func main() {
	// can hold a single value - unbuffered channel
	var myChannel = make(chan int)
	var my_other_channel = make(chan int)
	go process(myChannel)
	for i := range myChannel {
		fmt.Println("Received: ", i)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}
