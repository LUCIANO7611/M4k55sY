// 代码生成时间: 2025-08-25 04:06:02
package main

import (
    "crypto/rand"
    "encoding/binary"
    "net/http"
    "github.com/gin-gonic/gin"
    "math/big"
)

// RandomNumberResponse is the structure of the response for the random number endpoint.
type RandomNumberResponse struct {
    Number int64 `json:"number"`
}

// randomNumberHandler is a Gin handler function that generates a random number and returns it in the response.
func randomNumberHandler(c *gin.Context) {
    var numInt64 int64
    randNum, err := rand.Int(rand.Reader, big.NewInt(100000)) // Generate a random number between 0 and 100000
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to generate random number.",
        })
        return
    }
    numInt64 = randNum.Int64() // Convert big.Int to int64
    
    c.JSON(http.StatusOK, RandomNumberResponse{Number: numInt64})
}

func main() {
    router := gin.Default() // Create a new Gin router with default middlewares.
    router.GET("/random", randomNumberHandler) // Register the random number handler.
    // Start the server on port 8080.
    router.Run(":8080")
}
