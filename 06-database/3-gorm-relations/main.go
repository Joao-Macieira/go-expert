package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	category := Category{Name: "Eletronicos"}
	db.Create(&category)
	db.Create(&Product{
		Name:       "Mouse",
		Price:      50.00,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "12345",
		ProductID: 1,
	})

	var products []Product

	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	}
}
