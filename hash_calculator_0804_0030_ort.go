// 代码生成时间: 2025-08-04 00:30:54
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

// HashCalculatorHandler 处理器函数，用于计算哈希值
func HashCalculatorHandler(c *gin.Context) {
    input := c.PostForm("input")
    if input == "" {
        // 如果输入为空，则返回错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Input cannot be empty",
        })
        return
    }

    // 计算SHA-256哈希值
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // 返回哈希值结果
    c.JSON(http.StatusOK, gin.H{
        "hash": hashString,
    })
}

func main() {
    router := gin.Default()

    // 注册处理器和路径
    router.POST("/hash", HashCalculatorHandler)

    // 启动服务
    log.Fatal(router.Run(":8080"))
}
