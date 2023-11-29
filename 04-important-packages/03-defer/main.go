package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First line")
	fmt.Println("Second line")
	fmt.Println("Third line")
}
