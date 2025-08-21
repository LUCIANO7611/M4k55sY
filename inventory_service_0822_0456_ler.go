// 代码生成时间: 2025-08-22 04:56:04
 * This service provides RESTful API endpoints for managing inventory items.
 */

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Quantity int `json:"quantity"`
}

// inventoryStorage is a simple in-memory storage for inventory items.
var inventoryStorage = make(map[string]InventoryItem)

// setupRouter sets up the Gin router with the necessary routes and middleware.
func setupRouter() *gin.Engine {
    router := gin.Default()

    // Middleware for logging requests.
    router.Use(gin.Logger())

    // Middleware for recovering from panics.
    router.Use(gin.Recovery())

    // Inventory endpoint
    router.GET("/inventory", GetAllItems)
    router.POST("/inventory", AddItem)
    router.PUT("/inventory/:id", UpdateItem)
    router.DELETE("/inventory/:id", RemoveItem)

    return router
}

// GetAllItems retrieves all inventory items.
func GetAllItems(c *gin.Context) {
    var items []InventoryItem
    for _, item := range inventoryStorage {
        items = append(items, item)
    }
    c.JSON(http.StatusOK, items)
}

// AddItem adds a new inventory item.
func AddItem(c *gin.Context) {
    var newItem InventoryItem
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    inventoryStorage[newItem.ID] = newItem
    c.JSON(http.StatusCreated, newItem)
}

// UpdateItem updates an existing inventory item.
func UpdateItem(c *gin.Context) {
    var updatedItem InventoryItem
    id := c.Param("id")
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    if _, exists := inventoryStorage[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "item not found",
        })
        return
    }
    inventoryStorage[id] = updatedItem
    c.JSON(http.StatusOK, updatedItem)
}

// RemoveItem removes an inventory item by ID.
func RemoveItem(c *gin.Context) {
    id := c.Param("id")
    if _, exists := inventoryStorage[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "item not found",
        })
        return
    }
    delete(inventoryStorage, id)
    c.JSON(http.StatusOK, gin.H{
        "message": "item removed",
    })
}

func main() {
    router := setupRouter()
    router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
