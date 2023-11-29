package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully! Size: %d bytes\n", size)

	f.Close()

	//read file

	file, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// read file with streaming encoding

	fileBuffer, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileBuffer)
	bytesPerRow := 3
	buffer := make([]byte, bytesPerRow)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			fmt.Println("Finished")
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}
