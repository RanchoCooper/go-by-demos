package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 初始化引擎
	engine := gin.Default()

	// 注册路由和处理函数
	engine.Any("/", WebRoot)
	// 其他注册方式
	engine.Handle("GET", "/get", WebRoot)
	engine.GET("/someGet", WebRoot)
	engine.POST("/somePost", WebRoot)
	engine.PUT("/somePut", WebRoot)
	engine.DELETE("/someDelete", WebRoot)
	engine.PATCH("/somePatch", WebRoot)
	engine.HEAD("/someHead", WebRoot)
	engine.OPTIONS("/someOptions", WebRoot)

	// 动态路由
	engine.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello, %s", name)
	})

	// 路由分组
	v1 := engine.Group("/v1")
	{
		v1.POST("/login", WebRoot)
		v1.POST("/logout", WebRoot)
	}



	// 绑定端口，默认8080
	_ = engine.Run(":9205")

}

func WebRoot(context *gin.Context) {
	// 获取查询参数
	context.DefaultQuery("firstname", "Guest")
	context.Query("name")
	// 获取表单数据
	context.DefaultPostForm("msg", "something")
	context.PostForm("message")

	// 获取文件
	file, _ := context.FormFile("file")
	log.Println(file.Filename)

	context.String(http.StatusOK, "hello, world")
}

// path中的参数
func main2() {

}
