package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{} {
			"lang": "Go语言",
			"tag": "<br>",
		}

		context.AsciiJSON(http.StatusOK, data)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8090")
}
