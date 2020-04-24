package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
	Age int `gorm:"NOT NULL"`
	Earn int32 `gorm:"NOT NULL"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.LogMode(true)

	// migrate table and supplement new columns only, won't change any data
	db.AutoMigrate(&User{Name: "rancho"})

	db.Create(&User{Name: "rancho", Age: 25})
	var u User
	db.Find(&u, "name", "rancho")

	db.Model(&u).Update("name", "rancho", "age", 18)
}
