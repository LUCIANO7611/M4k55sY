// 代码生成时间: 2025-08-23 22:59:24
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

// 购物车项目
type CartItem struct {
    ProductID string `json:"product_id"`
    Quantity  int    `json:"quantity"`
}

// 购物车
type ShoppingCart struct {
    Items map[string]CartItem
}

// 添加商品到购物车
func AddToCart(c *gin.Context) {
    var item CartItem
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request data",
        })
        return
    }
    cart, exists := c.Get("cart")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Shopping cart not initialized",
        })
        return
    }
    cart.(*ShoppingCart).Items[item.ProductID] = item
    c.JSON(http.StatusOK, gin.H{
        "message": "Item added to cart",
    })
}

// 获取购物车内容
func GetCart(c *gin.Context) {
    cart, exists := c.Get("cart")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Shopping cart not initialized",
        })
        return
    }
    c.JSON(http.StatusOK, cart.(*ShoppingCart).Items)
}

// 初始化购物车中间件
func InitCartMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        cart := &ShoppingCart{Items: make(map[string]CartItem)}
        c.Set("cart", cart)
        c.Next()
    }
}

func main() {
    r := gin.Default()
    r.Use(InitCartMiddleware())

    // 添加商品到购物车
    r.POST("/add", AddToCart)

    // 获取购物车内容
    r.GET("/cart", GetCart)

    r.Run() // 默认在 8080 端口启动
}
