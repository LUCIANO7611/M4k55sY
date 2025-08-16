// 代码生成时间: 2025-08-17 02:09:03
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TestData 生成的测试数据结构
type TestData struct {
	ID        string    "json:"id""
	Timestamp time.Time "json:"timestamp""
	Value     string    "json:"value""
}

// generateTestDataHandler 处理生成测试数据的请求
func generateTestDataHandler(c *gin.Context) {
	// 日志记录请求开始时间
	startTime := time.Now()
	log.Printf("Generating test data at: %v", startTime)

	// 创建测试数据
	data := TestData{
		ID:        fmt.Sprintf("ID-%d", time.Now().UnixNano()),
		Timestamp: startTime,
		Value:     "Sample Value",
	}

	// 将测试数据写入响应体
	c.JSON(http.StatusOK, data)
}

// main 函数设置路由并启动Gin服务器
func main() {
	// 实例化Gin引擎
	r := gin.Default()

	// 设置日志中间件
	r.Use(gin.Logger())

	// 设置恢复中间件，用于错误恢复
	r.Use(gin.Recovery())

	// 设置测试数据生成器路由
	r.GET("/test-data", generateTestDataHandler)

	// 启动服务器
	log.Printf("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
