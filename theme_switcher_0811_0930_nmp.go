// 代码生成时间: 2025-08-11 09:30:41
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// ThemeSwitcherHandler 是主题切换的处理器
func ThemeSwitcherHandler(c *gin.Context) {
    // 获取主题名称作为URL参数
    theme := c.Param("theme")

    // 检查主题是否有效
    validThemes := []string{"light", "dark"}
    if contains(validThemes, theme) {
        // 设置主题到响应Cookie
        c.SetCookie("theme", theme, 86400, "/", "", false, true)
        c.JSON(http.StatusOK, gin.H{
            "message": "Theme switched to " + theme,
        })
    } else {
        // 如果主题无效，返回错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid theme",
        })
    }
}

// contains 检查slice中是否包含指定元素
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func main() {
    r := gin.Default()

    // 定义主题切换路由
    r.GET("/theme/:theme", ThemeSwitcherHandler)

    // 启动服务器
    r.Run(":8080")
}
