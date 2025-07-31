// 代码生成时间: 2025-07-31 14:27:13
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ThemeSwitcherHandler is a handler function that switches the theme.
// It takes a Gin context and a theme name as parameters.
// It returns an error if theme switching fails.
func ThemeSwitcherHandler(c *gin.Context, theme string) error {
    // Check if the theme is valid
    validThemes := []string{"light", "dark"}
    if !isThemeValid(theme, validThemes) {
        return c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid theme specified.",
        })
    }
    
    // Simulate setting the theme in the user's session or cookie
    // For simplicity, this example will just set a cookie
    c.SetCookie("theme", theme, 3600, "/", "", false, true)
    
    // Return a success message with the chosen theme
    return c.JSON(http.StatusOK, gin.H{
        "message": "Theme switched successfully.",
        "theme": theme,
    })
}

// isThemeValid checks if a given theme is present in the list of valid themes.
func isThemeValid(theme string, validThemes []string) bool {
    for _, v := range validThemes {
        if v == theme {
            return true
        }
    }
    return false
}

func main() {
    router := gin.Default()
    
    // Define the route for theme switching
    router.POST("/switch-theme", func(c *gin.Context) {
        theme := c.PostForm("theme")
        if err := ThemeSwitcherHandler(c, theme); err != nil {
            // Handle any errors encountered during theme switching
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "error": "An error occurred while switching the theme.",
            })
            return
        }
    })
    
    // Start the server
    router.Run(":8080")
}
