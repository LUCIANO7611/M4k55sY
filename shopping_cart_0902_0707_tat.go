// 代码生成时间: 2025-09-02 07:07:49
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ShoppingCart represents a shopping cart with a list of items.
type ShoppingCart struct {
    Items []CartItem `json:"items"`
}

// CartItem represents an item in the shopping cart.
type CartItem struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Quantity int    `json:"quantity"`
}

func main() {
    r := gin.Default()

    // Define the shopping cart
    var cart ShoppingCart

    // Endpoint to add an item to the cart
    r.POST("/add", func(c *gin.Context) {
        var item CartItem
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        // Check for item existence and add or update it
        for i, existingItem := range cart.Items {
            if existingItem.ID == item.ID {
                cart.Items[i].Quantity += item.Quantity
                c.JSON(http.StatusOK, gin.H{
                    "message": "Item updated in cart",
                "cart": cart,
                })
                return
            }
        }
        cart.Items = append(cart.Items, item)
        c.JSON(http.StatusOK, gin.H{
            "message": "Item added to cart",
            "cart": cart,
        })
    })

    // Endpoint to remove an item from the cart
    r.DELETE("/remove/:id", func(c *gin.Context) {
        id := c.Param("id")
        var found bool
        for i, item := range cart.Items {
            if item.ID == id {
                found = true
                cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
                break
            }
        }
        if !found {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Item not found in cart",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Item removed from cart",
            "cart": cart,
        })
    })

    // Start the server
    r.Run()
}
