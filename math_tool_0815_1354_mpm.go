// 代码生成时间: 2025-08-15 13:54:22
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// CustomError 定义了自定义错误结构体
type CustomError struct {
    Title       string `json:"title"`
    Status      string `json:"status"`
    Detail      string `json:"detail"`
    Instance    string `json:"instance"`
}

// ErrorResponse 定义了返回错误信息的结构体
type ErrorResponse struct {
    Error CustomError `json:"error"`
}

// MathCalculator 定义了数学计算的结构体
type MathCalculator struct{}

// Add 执行加法运算
func (c *MathCalculator) Add(ctx *gin.Context) {
    var req struct {
        A float64 `json:"a" binding:"required,number"`
        B float64 `json:"b" binding:"required,number"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: CustomError{
            Title:   "Invalid request",
            Status:  "400",
            Detail:  "Invalid input for add operation",
            Instance: ctx.Request.URL.String(),
        }})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "result": req.A + req.B,
    })
}

// Subtract 执行减法运算
func (c *MathCalculator) Subtract(ctx *gin.Context) {
    var req struct {
        A float64 `json:"a" binding:"required,number"`
        B float64 `json:"b" binding:"required,number"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: CustomError{
            Title:   "Invalid request",
            Status:  "400",
            Detail:  "Invalid input for subtract operation",
            Instance: ctx.Request.URL.String(),
        }})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "result": req.A - req.B,
    })
}

func main() {
    router := gin.Default()

    // 使用中间件记录请求日志
    router.Use(gin.Logger())

    // 使用中间件恢复任何发生的panic，并将堆栈信息返回给客户端
    router.Use(gin.Recovery())

    calc := MathCalculator{}

    // 添加加法和减法的路由
    router.POST("/add", calc.Add)
    router.POST("/subtract", calc.Subtract)

    // 启动服务器
    router.Run(":8080")
}
