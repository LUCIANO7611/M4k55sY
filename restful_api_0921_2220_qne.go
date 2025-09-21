// 代码生成时间: 2025-09-21 22:20:01
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse 定义了一个错误响应的结构体
type ErrorResponse struct {
    Error string `json:"error"`
}

// Item 定义了一个项目的结构体
type Item struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

// getAllItemsHandler 处理获取所有项目的请求
func getAllItemsHandler(c *gin.Context) {
    items := []Item{
        {ID: 1, Name: "item1"},
        {ID: 2, Name: "item2"},
    }
    c.JSON(http.StatusOK, items)
}

// getItemHandler 处理根据ID获取项目的请求
func getItemHandler(c *gin.Context) {
    id := c.Param("id")
    // 这里可以添加更多的错误处理和验证逻辑
    c.JSON(http.StatusOK, Item{ID: uint(idInt), Name: "item" + id})
}

// createItemHandler 处理创建新项目的请求
func createItemHandler(c *gin.Context) {
    var item Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }
    // 这里可以添加逻辑来处理创建项目，例如存储到数据库
    c.JSON(http.StatusCreated, item)
}

func main() {
    r := gin.Default()

    // 使用中间件记录请求日志
    r.Use(gin.Logger())
    // 使用中间件恢复处理任何 panic，这样服务器不会崩溃并返回 500 错误
    r.Use(gin.Recovery())

    // 定义路由
    r.GET("/items", getAllItemsHandler)
    r.GET("/items/:id", getItemHandler)
    r.POST("/items", createItemHandler)

    // 启动服务
    fmt.Println("Server started at :8080")
    r.Run(":8080")
}
