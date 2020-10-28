package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移schema
	_ = db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	// read
	var product Product
	// 根据主键查找
	db.First(&product, 1)
	// 根据条件查找
	db.First(&product, "code = ?", "D42")

	// update
	db.Model(&product).Update("Price", 200)
	// 仅更新非零值字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// delete
	db.Delete(&product, 1)
}
