package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type QuotationModel struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Quotation struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	server := http.NewServeMux()
	server.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
		defer cancel()

		endpoint := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
		req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		var quotation Quotation
		err = json.Unmarshal(body, &quotation)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		cotationModel := QuotationModel{
			Code:       quotation.USDBRL.Code,
			Codein:     quotation.USDBRL.Codein,
			Name:       quotation.USDBRL.Name,
			High:       quotation.USDBRL.High,
			Low:        quotation.USDBRL.Low,
			VarBid:     quotation.USDBRL.VarBid,
			PctChange:  quotation.USDBRL.PctChange,
			Bid:        quotation.USDBRL.Bid,
			Ask:        quotation.USDBRL.Ask,
			Timestamp:  quotation.USDBRL.Timestamp,
			CreateDate: quotation.USDBRL.CreateDate,
		}

		err = SaveQuotation(db, &cotationModel)

		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cotationModel)
	})

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}

func SaveQuotation(db *sql.DB, quotation *QuotationModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db.Exec("CREATE TABLE IF NOT EXISTS quotations (code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, createDate TEXT)")

	return db.QueryRowContext(ctx, "INSERT INTO quotations VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING bid",
		quotation.Code,
		quotation.Codein,
		quotation.Name,
		quotation.High,
		quotation.Low,
		quotation.VarBid,
		quotation.PctChange,
		quotation.Bid,
		quotation.Ask,
		quotation.Timestamp,
		quotation.CreateDate,
	).Scan(&quotation.Bid)
}
