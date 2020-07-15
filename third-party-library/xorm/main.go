package main

import (
	"fmt"
	"git.llsapp.com/awesome/log"
	"github.com/go-xorm/xorm"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64
	Name     string
	Salt     string
	Age      int
	Password string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func main() {
	// 创建一个引擎
	engine, err := xorm.NewEngine("mysql", "root@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	// 输出执行的SQL
	engine.ShowSQL(true)

	// 根据传入的struct结构自动创建对应的表
	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	// 使用主键查询
	user1 := &User{}
	has, err := engine.ID(1).Get(user1)
	if has {
	}
	// 使用where子句查询
	user2 := &User{}
	has, err = engine.Where("name=?", "rancho").Get(user2)
	if has {
	}
	// 使用对象的非空字段查询
	user3 := &User{Id: 5}
	has, err = engine.Get(user3)

	// 只返回指定的列
	user4 := &User{}
	has, err = engine.ID(1).Cols("id", "name", "age").Get(user4)

	// 忽略部分列
	user5 := &User{}
	has, err = engine.Omit("created", "updated").Get(user5)

	// 查询符合条件的记录是否存在
	has, _ = engine.ID(1).Exist(user1)

	// 查询所有符合条件的记录，Get只查一条(LIMIT 1)
	users := make([]User, 1)
	err = engine.Where("age > ? and age < ?", 12, 30).Find(&users)

	// Iterate 回调处理每条记录
	engine.Where("age > ? and age < ?", 12, 30).
		Iterate(&User{}, func(i int, bean interface{}) error {
			fmt.Printf("user%d:%v\n", i, bean.(*User))
			return nil
		})
}
