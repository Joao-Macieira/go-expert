package main

import "fmt"

func receive(name string, hello chan<- string) { // only receive information in this channel (arrow on right)
	hello <- name
}

func read(data <-chan string) { // only emptied the channel (arrow on left)
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}
