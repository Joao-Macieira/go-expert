package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
