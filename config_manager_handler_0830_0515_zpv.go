// 代码生成时间: 2025-08-30 05:15:03
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "github.com/gin-gonic/gin"
)

// ConfigManagerHandler 处理配置文件管理
func ConfigManagerHandler(c *gin.Context) {
    // 获取配置文件路径参数
    filePath := c.DefaultQuery("path", "")

    // 检查文件路径是否提供
    if filePath == "" {
        c.JSON(400, gin.H{
            "error": "File path is required",
        })
        return
    }

    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(404, gin.H{
            "error": "File not found",
        })
        return
    }

    // 读取文件内容
    file, err := os.ReadFile(filePath)
    if err != nil {
        c.JSON(500, gin.H{
            "error": "Failed to read file",
        })
        return
    }

    // 发送文件内容作为响应
    c.Data(200, "application/json", file)
}

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine) {
    // 添加配置文件管理器路由
    router.GET("/config", ConfigManagerHandler)
}

// main 函数设置Gin并启动服务器
func main() {
    // 创建Gin路由器
    router := gin.Default()

    // 设置路由
    SetupRoutes(router)

    // 启动服务器
    log.Printf("Server started at :8080")
    if err := router.Run(":8080"); err != nil {
        fmt.Printf("Failed to start server: %v", err)
    }
}
