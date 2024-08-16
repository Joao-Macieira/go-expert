package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // empty

	// Thread 2
	go func() {
		channel <- "Hello world" // full
	}()

	// Thread 1
	msg := <-channel // emptied
	fmt.Println(msg)
}
