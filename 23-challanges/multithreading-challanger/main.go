package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ViaCEPResult struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

type BrasilAPIResult struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

// Função para buscar o endereço na API ViaCEP
func fetchViaCEP(cep string, ch chan<- interface{}) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Erro na API ViaCEP: %v", err)
		return
	}
	defer response.Body.Close()

	var result ViaCEPResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		ch <- fmt.Sprintf("Erro ao decodificar resposta ViaCEP: %v", err)
		return
	}
	ch <- result
}

// Função para buscar o endereço na API BrasilAPI
func fetchBrasilAPI(cep string, ch chan<- interface{}) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Erro na API BrasilAPI: %v", err)
		return
	}
	defer response.Body.Close()

	var result BrasilAPIResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		ch <- fmt.Sprintf("Erro ao decodificar resposta BrasilAPI: %v", err)
		return
	}
	ch <- result
}

func main() {
	cep := "01153000"
	ch := make(chan interface{})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go fetchViaCEP(cep, ch)
	go fetchBrasilAPI(cep, ch)

	select {
	case result := <-ch:
		switch res := result.(type) {
		case ViaCEPResult:
			fmt.Println("API ViaCEP result:")
			fmt.Printf("CEP: %s, Logradouro: %s, Bairro: %s, Cidade: %s, UF: %s\n",
				res.Cep, res.Logradouro, res.Bairro, res.Localidade, res.Uf)
		case BrasilAPIResult:
			fmt.Println("API BrasilAPI result:")
			fmt.Printf("CEP: %s, Logradouro: %s, Bairro: %s, Cidade: %s, UF: %s\n",
				res.Cep, res.Street, res.Neighborhood, res.City, res.State)
		default:
			fmt.Println("Erroe: Invalid type received")
		}
	case <-ctx.Done():
		fmt.Println("Error: Request timed out")
	}
}
