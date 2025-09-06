// 代码生成时间: 2025-09-07 06:40:36
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// setupRouter sets up the Gin router with routes and middlewares.
func setupRouter() *gin.Engine {
    r := gin.Default()

    // Add middlewares if needed, for example, Logger and Recovery
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // Define a route for the home page with error handling
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to the responsive layout design page!",
        })
    })

    // Define error handling route
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "404 Not Found",
        })
    })

    return r
}

func main() {
    router := setupRouter()

    // Start the server on port 8080
    fmt.Println("Server started on port 8080")
    if err := router.Run(":8080"); err != nil {
        panic(fmt.Sprintf("Failed to start server: %v", err))
    }
}
