package main

import (
	server "Approval_dechuang/app/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	projectRoot, err := os.Getwd()
	//projectRoot = "E:\\Approval_dechuang"
	fmt.Println(projectRoot)
	if err != nil {
		panic(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.DELETE("/delete-files", server.DeleteHandlerWithDir(projectRoot))
	r.Run(":8080")
}
