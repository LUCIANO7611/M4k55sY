// 代码生成时间: 2025-09-09 04:48:53
package main

import (
    "net/http"
    "strings"
    "html"
    "github.com/gin-gonic/gin"
)

// XssMiddleware is a Gin middleware function that prevents XSS attacks by
// sanitizing the input to remove any HTML tags.
func XssMiddleware(c *gin.Context) {
    // Get the raw query string
    rawQuery := c.Request.URL.RawQuery
    // Remove HTML tags from the query string to prevent XSS attacks
    sanitizedQuery := html.EscapeString(rawQuery)
    // Replace the original query string with the sanitized one
    c.Request.URL.RawQuery = sanitizedQuery
    // Proceed to the next middleware
    c.Next()
}

// ErrorHandler is a Gin middleware function that handles errors and
// returns a JSON response with an error message.
func ErrorHandler(c *gin.Context) {
    c.Next()
    // If this request has an error recorded, handle it
    if len(c.Errors) > 0 {
        // Get the first error (there should only be one)
        err := c.Errors[0].Err
        // Log the error (not implemented here)
        // log.Printf("Error: %v", err)
        // Respond with a JSON error message
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
    }
}

func main() {
    r := gin.Default()
    // Add the XssMiddleware to the router
    r.Use(XssMiddleware)
    // Add the ErrorHandler middleware
    r.Use(ErrorHandler)

    // Define a simple route to demonstrate the middleware
    r.GET("/", func(c *gin.Context) {
        // Retrieve the sanitized query from the context
        sanitizedQuery := c.Request.URL.RawQuery
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, your query is sanitized!",
            "sanitizedQuery": sanitizedQuery,
        })
    })

    // Start the server
    r.Run(":8080") // Listen and serve on 0.0.0.0:8080
}
