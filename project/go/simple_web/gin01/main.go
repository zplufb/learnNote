/**
 * @Author: FB
 * @Description: 
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/1/2 10:28
 */

package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	r := gin.Default()         //获得路由实例

	//添加中间件
	r.Use(Middleware)

	r.Static("/html", "./gin01/web/page/html")

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	r.RunTLS(":443", "./server.pem", "./server.key")

	//监听端口
	//r.Run(":9301")
	//http.ListenAndServe(":9301", r)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}
