package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Microsecond,
	}

	response, err := client.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
