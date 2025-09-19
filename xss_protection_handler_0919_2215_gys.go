// 代码生成时间: 2025-09-19 22:15:31
package main

import (
    "net/http"
    "strings"
    "html"
    "github.com/gin-gonic/gin"
)

// XSSProtectionHandler is a Gin middleware that provides basic XSS protection by sanitizing the input.
func XSSProtectionHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Iterate over all form data
        formData := c.Request.PostForm
        for key, values := range formData {
            for i, value := range values {
                // Sanitize the value to prevent XSS attacks
                sanitized := html.EscapeString(value)
                // Update the form data with the sanitized value
                values[i] = sanitized
            }
        }

        // Proceed to the next middleware or handler
        c.Next()
    }
}

// ErrorHandler handles errors by returning a JSON response with an error code and message.
func ErrorHandler(c *gin.Context) {
    c.JSON(http.StatusInternalServerError, gin.H{
        "error": "Internal Server Error",
    })
}

func main() {
    r := gin.Default()

    // Register the XSS protection middleware
    r.Use(XSSProtectionHandler())

    // Register a sample route to demonstrate the XSS protection
    r.GET("/xss", func(c *gin.Context) {
        userInput := c.DefaultQuery("input", "")
        c.HTML(http.StatusOK, "xss.html", gin.H{
            "userInput": userInput,
        })
    })

    // Register the error handler for all routes
    r.NoRoute(ErrorHandler)

    // Start the server
    r.Run()
}

// Note: This example assumes you have an 'xss.html' template file for rendering the user input.
// Make sure to create the template file in the 'templates' directory and configure Gin's template
// settings accordingly.
