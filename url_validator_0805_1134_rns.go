// 代码生成时间: 2025-08-05 11:34:56
package main

import (
    "net/url"
    "strings"
    "github.com/gin-gonic/gin"
)

// validateURL checks if the provided URL is valid and returns a boolean.
func validateURL(c *gin.Context) {
    providedURL := c.Query("url")
    if providedURL == "" {
        c.JSON(400, gin.H{
            "error": "URL parameter is missing",
        })
        return
    }
    
    // Parse the URL to check its validity.
    u, err := url.ParseRequestURI(providedURL)
    if err != nil {
        c.JSON(400, gin.H{
            "error": "Invalid URL format",
        })
        return
    }
    
    // Check if the scheme is http or https as a basic validation.
    if !strings.HasPrefix(u.Scheme, "http") {
        c.JSON(400, gin.H{
            "error": "URL scheme must be HTTP or HTTPS",
        })
        return
    }
    
    c.JSON(200, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    r := gin.Default()
    
    // Register the URL validation handler.
    r.GET("/validate", validateURL)
    
    // Start the HTTP server.
    r.Run() // listen and serve on 0.0.0.0:8080
}
