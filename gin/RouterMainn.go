package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("Request path: ", context.FullPath())

		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("hello, " + name))
	})

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println("Request path: ", context.FullPath())

		username, exists := context.GetPostForm("username")
		if exists {
			fmt.Println(username)
		}
		password, exists := context.GetPostForm("password")
		if exists {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("hello, " + username))
	})

	engine.Run(":9090")
}
