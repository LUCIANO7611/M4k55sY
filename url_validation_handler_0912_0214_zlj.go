// 代码生成时间: 2025-09-12 02:14:10
package main

import (
    "net/url"
    "strings"
    "github.com/gin-gonic/gin"
    "net/http"
)

// URLValidationHandler 验证URL链接有效性的处理器
func URLValidationHandler(c *gin.Context) {
    // 从请求中获取URL参数
    reqURL := c.Query("url")

    // 检查URL参数是否存在
    if reqURL == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL parameter is required",
        })
        return
    }

    // 尝试解析URL
    parsedURL, err := url.ParseRequestURI(reqURL)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL format",
        })
        return
    }

    // 检查URL是否有有效的Scheme和Host
    if parsedURL.Scheme == "" || parsedURL.Host == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL must have a valid scheme and host",
        })
        return
    }

    // 如果URL有效，返回成功状态
    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

// main 函数，设置Gin路由器并注册URL验证处理器
func main() {
    router := gin.Default()

    // 注册URL验证处理器
    router.GET("/check-url", URLValidationHandler)

    // 启动服务器
    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
