// 代码生成时间: 2025-08-29 09:28:16
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

// HashCalculatorHandler 处理哈希值计算请求
func HashCalculatorHandler(c *gin.Context) {
    // 从请求中获取要计算哈希的字符串
    input := c.PostForm("input")

    // 检查输入是否为空
    if input == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Input is required",
        })
        return
    }

    // 计算MD5哈希值
    hash := md5.Sum([]byte(input))
    hashStr := hex.EncodeToString(hash[:])

    // 返回哈希值
    c.JSON(http.StatusOK, gin.H{
        "hash": hashStr,
    })
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()

    // 注册哈希值计算处理程序
    router.POST("/hash", HashCalculatorHandler)

    // 设置端口号并启动服务器
    log.Fatal(router.Run(":8080"))
}
