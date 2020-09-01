package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// create gin engine with default configs
	engine := gin.Default()

	// add '/hello' view function with handle function
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("request path: ", context.FullPath())
		context.Writer.Write([]byte("hello, gin\n"))
	})

	// run server with specific port
	if err := engine.Run(":9090"); err != nil {
		log.Fatal(err.Error())
	}
}

