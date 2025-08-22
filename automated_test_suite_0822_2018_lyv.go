// 代码生成时间: 2025-08-22 20:18:24
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse 结构体定义用于返回错误信息
type ErrorResponse struct {
    Error string `json:"error"`
}

// 定义测试组数据
var testData = []struct {
	Input  string
	Expect int
}{
	{Input: "2+2", Expect: 4},
	{Input: "3+3", Expect: 6},
	{Input: "4+4", Expect: 8},
}

// TestHandler 处理器返回测试结果
func TestHandler(c *gin.Context) {
	for _, data := range testData {
		result, err := evaluateExpression(data.Input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		if result != data.Expect {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Test failed for input: " + data.Input})
			return
		}
	}

c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": "All tests passed",
})
}

// evaluateExpression 处理字符串表达式并返回结果
func evaluateExpression(expression string) (int, error) {
	// 这里只是一个示例，实际情况可能需要使用更复杂的解析器
	// 假设表达式格式为 "数字+数字"
	parts := strings.Split(expression, "+")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid expression format")
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	right, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return left + right, nil
}

func main() {
	router := gin.Default()

	// 使用中间件
	router.Use(gin.Recovery())

	// 注册测试处理器
	router.GET("/test", TestHandler)

	// 启动服务
	router.Run(":8080")
}
