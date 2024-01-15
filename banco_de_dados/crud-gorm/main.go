package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Catergory struct {
	Name string
	gorm.Model
}

type Product struct {
	Name        string
	Price       float64
	CatergoryID uint
	Catergory   Catergory
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name:  "notebook",
	// 	Price: 2000.9,
	// })

	// // create batch
	// products := []Product{
	// 	{Name: "mouse", Price: 20.9},
	// 	{Name: "keyboard", Price: 50.9},
	// }
	// db.Create(&products)

	// select one

	// var product Product
	// // db.First(&product, 2)

	// db.First(&product, "name = ?", "mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)

	// for _, p := range products {
	// 	println(p.Name)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 1000).Find(&products)
	// for _, p := range products {
	// 	println(p.Name)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "teste"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// println(p2.Name)
	// db.Delete(&p2)
}
