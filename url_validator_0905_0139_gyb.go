// 代码生成时间: 2025-09-05 01:39:38
package main

import (
    "net/http"
    "regexp"

    "github.com/gin-gonic/gin"
)

// URLValidatorMiddleware 是一个Gin中间件，用于验证URL链接的有效性
func URLValidatorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        url := c.Request.URL.String()
        // 编译URL正则表达式
        re := regexp.MustCompile(`^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
        // 检查URL是否符合正则表达式
        if !re.MatchString(url) {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid URL"
            })
            c.Abort() // 停止处理请求
            return
        }
        // 继续处理请求
        c.Next()
    }
}

// setupRoutes 设置路由和中间件
func setupRoutes(r *gin.Engine) {
    // 使用URL验证中间件
    r.Use(URLValidatorMiddleware())
    r.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "URL is valid"
        })
    })
}

func main() {
    r := gin.Default()
    setupRoutes(r)
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}

// 请注意，这个正则表达式是一个非常基础的URL验证器，可能不适用于所有情况。
// 复杂URL的验证可能需要更精细的正则表达式或专门的库。