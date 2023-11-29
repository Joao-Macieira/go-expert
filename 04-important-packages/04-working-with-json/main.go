package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"`
}

func main() {
	account := Account{Number: 1, Balance: 100}

	result, err := json.Marshal(account)

	c := reflect.TypeOf(account).Field(0).Tag
	fmt.Printf("%s\n\n", c)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	err = json.NewEncoder(os.Stdout).Encode(account)

	if err != nil {
		panic(err)
	}

	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(account)

	fakeJson := []byte(`{"number":2,"balance":300}`)
	var accountX Account

	err = json.Unmarshal(fakeJson, &accountX)

	if err != nil {
		panic(err)
	}

	fmt.Println(accountX)
}
