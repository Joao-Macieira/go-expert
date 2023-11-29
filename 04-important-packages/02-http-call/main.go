package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	result, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	request.Body.Close()
}
