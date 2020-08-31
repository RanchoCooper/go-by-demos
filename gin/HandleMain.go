package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// create gin engine with default configs
	engine := gin.Default()

	// http://localhost:9090/hello?name=rancho
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println("Request path: ", path)

		// get query parameters
		name := context.DefaultQuery("name", "rancho")
		fmt.Println("receive name: ", name)

		context.Writer.Write([]byte("hello, " + name))
	})

	// POST http://localhost:9090/login
	engine.Handle("POST", "/login", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println("Request path: ", path)

		username := context.PostForm("username")
		password := context.PostForm("password")


		context.Writer.Write([]byte("welcome, " + username + " you have login with password: " + password))
	})

	engine.Run(":9090")
}

