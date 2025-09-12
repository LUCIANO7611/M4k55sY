// 代码生成时间: 2025-09-13 05:17:39
package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Item represents an inventory item
type Item struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
}

// InventoryService handles inventory operations
type InventoryService struct {
    items []Item
}

// NewInventoryService creates a new InventoryService
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: []Item{},
    }
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid item data'})
        return
    }
    s.items = append(s.items, newItem)
    c.JSON(http.StatusOK, newItem)
}

// GetItems returns the list of items in the inventory
func (s *InventoryService) GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, s.items)
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(c *gin.Context) {
    id := c.Param("id")
    var updatedItem Item
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid item data'})
        return
    }
    for i, item := range s.items {
        if item.ID == id {
            s.items[i] = updatedItem
            c.JSON(http.StatusOK, s.items[i])
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{'error': 'Item not found'})
}

// DeleteItem removes an item from the inventory
func (s *InventoryService) DeleteItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range s.items {
        if item.ID == id {
            copy(s.items[i:], s.items[i+1:])
            s.items[len(s.items)-1] = Item{}
            s.items = s.items[:len(s.items)-1]
            c.JSON(http.StatusOK, gin.H{"success": "Item deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{'error': 'Item not found'})
}

func main() {
    r := gin.Default()

    // Initialize InventoryService
    service := NewInventoryService()

    // Use middleware
    r.Use(gin.Logger(), gin.Recovery())

    // Routes
    r.POST("/items", service.AddItem)
    r.GET("/items", service.GetItems)
    r.PUT("/items/:id", service.UpdateItem)
    r.DELETE("/items/:id", service.DeleteItem)

    // Start server
    r.Run() // listening and serving on 0.0.0.0:8080
}
