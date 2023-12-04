package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "joão"}`))

	response, err := client.Post("http://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil)
}
