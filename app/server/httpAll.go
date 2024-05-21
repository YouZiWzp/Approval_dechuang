package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RemoveAll(dirPath string) error {
	//var volume string
	//if len(dirPath) >= 2 && dirPath[1] == ':' {
	//	volume := dirPath[:3]
	//	fmt.Println("盘符是:", volume)
	//} else {
	//	fmt.Println("这不是一个有效的盘符路径")
	//}
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("directory %s does not exist", dirPath)
	}

	// 删除目录及其内容
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}

	return nil
}
func DeleteHandlerWithDir(dirPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调用 removeAll 函数删除目录内容
		if err := RemoveAll(dirPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 发送包含 IP 地址的邮件（这里只是伪代码）
		// ...

		c.JSON(http.StatusOK, gin.H{"message": "Files deleted successfully"})
	}
}
