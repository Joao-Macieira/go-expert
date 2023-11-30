package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		request, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on make request: %v\n", err)
		}

		defer request.Body.Close()

		response, err := io.ReadAll(request.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on read response: %v\n", err)
		}

		var data ViaCEP

		err = json.Unmarshal(response, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on parse response: %v\n", err)
		}

		file, err := os.Create("city.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on create file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on write in file: %v\n", err)
		}

		fmt.Println("File created successfully")
		fmt.Println("CEP: ", data.Cep)
		fmt.Println("Localidade: ", data.Localidade)
		fmt.Println("UF: ", data.Uf)
	}
}
