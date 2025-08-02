// 代码生成时间: 2025-08-02 19:34:27
 * interactive_chart_handler.go
 * This file contains a Gin handler for generating interactive charts.
 */

package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ChartDataRequest is a struct to hold request data for chart generation
type ChartDataRequest struct {
    Type    string `form:"type" json:"type"`
    Options map[string]string `form:"options" json:"options"`
}

// ChartDataResponse is a struct to hold response data for chart generation
type ChartDataResponse struct {
    ChartType    string `json:"chartType"`
    ChartOptions string `json:"chartOptions"`
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery()) // Use Gin recovery middleware for panic recovery

    // Define a route for chart generation with POST method
    r.POST("/generateChart", generateChartHandler)
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Interactive Chart Generator API"
        })
    })

    // Start the server on port 8080
    r.Run(":8080")
}

// generateChartHandler is the handler function to generate interactive charts
func generateChartHandler(c *gin.Context) {
    var reqData ChartDataRequest
    if err := c.ShouldBindJSON(&reqData); err != nil {
        // Handle any JSON parsing errors
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request data"
        })
        return
    }

    // Generate chart data based on request
    chartResponse := ChartDataResponse{
        ChartType:    reqData.Type,
        ChartOptions: fmt.Sprintf("%+v", reqData.Options),
    }

    // Use chartResponse to generate chart (implementation not shown)
    // ... (chart generation logic)

    // Return the chart data in the response
    c.JSON(http.StatusOK, chartResponse)
}
