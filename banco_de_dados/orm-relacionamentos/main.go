package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Product struct {
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	Number    string
	ProductID uint
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// // create category
	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	// create product
	db.Create(&Product{
		Name:       "panela",
		Price:      2000.9,
		Categories: []Category{category, category2},
	})

	// // create serial number
	// db.Create(&SerialNumber{
	// 	Number:    "123456789",
	// 	ProductID: 1,
	// })

	// select all
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	// }

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		println("\ncategory", c.Name, "--")
		for _, p := range c.Products {
			print("/", p.Name)
		}
	}

}
