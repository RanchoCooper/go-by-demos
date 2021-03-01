package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func sqliteExample() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate schema
	_ = db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Code: "D42",
		Price: 100,
	})
	db.Commit()

	// read
	var product Product

	db.First(&product, "code = ?", "D42")
	fmt.Println(product)

	// update
	db.Model(&product).Update("Price", 200)
	fmt.Println(product)
	db.Model(&product).Updates(Product{Price: 400, Code: "F42"})
	fmt.Println(product)

	// delete
	db.Delete(&product, product.ID)
	db.Commit()
}

func mysqlExample()  {
	dsn := "root:rancho@tcp(127.0.0.1:3306)/test"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	fmt.Println(db)
}

func main()  {
	sqliteExample()
	mysqlExample()
}