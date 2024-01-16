package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 2999.90,
	// })

	// products := []Product{
	// 	{Name: "Playstation 5", Price: 4000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 129.00},
	// }
	// db.Create(&products)

	// var product Product
	// db.Find(&product, 4)
	// db.Find(&product, "name = ?", "keyboard")
	// fmt.Println(product)

	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// var p Product
	// db.Find(&p, 1)
	// p.Name = "Desktop"
	// db.Save(&p)

	// var searchedP Product
	// db.Find(&searchedP, 1)
	// fmt.Println(searchedP)
	// db.Delete(&searchedP)
}
