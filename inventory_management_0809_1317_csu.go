// 代码生成时间: 2025-08-09 13:17:00
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Count int    `json:"count"`
}

// inventory is a slice of InventoryItems
var inventory []InventoryItem

// InitializeInventory populates the inventory with sample data
func InitializeInventory() {
    inventory = []InventoryItem{{ID: 1, Name: "Item1", Count: 100},
                            {ID: 2, Name: "Item2", Count: 200}}
}

// GetInventoryItems returns all inventory items
func GetInventoryItems(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "inventory": inventory,
    })
}

// GetInventoryItemByID finds and returns the inventory item by ID
func GetInventoryItemByID(c *gin.Context) {
    id := c.Param("id")
    for _, item := range inventory {
        if item.ID == int(gin.Int64(id)) {
            c.JSON(http.StatusOK, item)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": "Inventory item not found",
    })
}

// AddInventoryItem adds a new item to the inventory
func AddInventoryItem(c *gin.Context) {
    var newItem InventoryItem
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    inventory = append(inventory, newItem)
    c.JSON(http.StatusOK, newItem)
}

// UpdateInventoryItem updates an existing inventory item
func UpdateInventoryItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range inventory {
        if item.ID == int(gin.Int64(id)) {
            var updatedItem InventoryItem
            if err := c.ShouldBindJSON(&updatedItem); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
            inventory[i] = updatedItem
            c.JSON(http.StatusOK, updatedItem)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": "Inventory item not found",
    })
}

// DeleteInventoryItem removes an inventory item by ID
func DeleteInventoryItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range inventory {
        if item.ID == int(gin.Int64(id)) {
            inventory = append(inventory[0:i], inventory[i+1:]...)
            c.JSON(http.StatusOK, gin.H{
                "message": "Inventory item deleted",
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{
        "error": "Inventory item not found",
    })
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery()) // Use Recovery middleware to handle panics

    InitializeInventory()

    r.GET("/inventory", GetInventoryItems)
    r.GET("/inventory/:id", GetInventoryItemByID)
    r.POST("/inventory", AddInventoryItem)
    r.PUT("/inventory/:id", UpdateInventoryItem)
    r.DELETE("/inventory/:id", DeleteInventoryItem)

    r.Run() // listen and serve on 0.0.0.0:8080
}
