// 代码生成时间: 2025-08-04 23:53:29
package main

import (
    "html"
    "net/http"
    "github.com/gin-gonic/gin"
)

// XSSFilterMiddleware 是一个中间件，用于防止XSS攻击
// 它将清理输入数据，以防止恶意脚本注入
func XSSFilterMiddleware(c *gin.Context) {
    // 从请求中获取所有参数
    parameters := c.Request.URL.Query()
    for key, value := range parameters {
        // 清除每个参数值中的HTML标签，防止XSS攻击
        cleanedValue := html.EscapeString(value[0])
        parameters.Set(key, cleanedValue)
    }
    // 更新URL中的查询参数
    c.Request.URL.RawQuery = parameters.Encode()
    c.Next()
}

// ErrorHandler 是一个自定义的错误处理函数
// 它将捕获并处理在请求处理过程中发生的错误
func ErrorHandler(c *gin.Context) {
    c.Next()
    if err, ok := c.Get(gin.ErrorKey); ok && err != nil {
        // 记录错误信息
        // log.Printf("Error: %v", err)
        // 响应错误信息给客户端
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal Server Error",
        })
    }
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()

    // 使用XSSFilterMiddleware中间件
    router.Use(XSSFilterMiddleware)

    // 使用自定义错误处理中间件
    router.Use(ErrorHandler)

    // 定义一个简单的GET路由，用于测试XSS攻击防护
    router.GET("/test", func(c *gin.Context) {
        // 获取请求参数
        input := c.Query("input")
        // 响应回客户端，显示输入的内容
        c.JSON(http.StatusOK, gin.H{
            "input": input,
        })
    })

    // 启动服务器
    router.Run(":8080")
}
