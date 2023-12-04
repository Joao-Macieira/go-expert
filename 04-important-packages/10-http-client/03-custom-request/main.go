package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.Client{}

	request, err := http.NewRequest("GET", "https://google.com", nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)

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
