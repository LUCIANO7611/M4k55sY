// 代码生成时间: 2025-09-14 14:47:30
package main

import (
    "excelize"
    "github.com/gin-gonic/gin"
    "net/http"
)

// ExcelGeneratorHandler 是一个 Gin handler，用于生成 Excel 文件
func ExcelGeneratorHandler(c *gin.Context) {
    title := c.Query("title")
    if title == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Title parameter is required",
        })
        return
    }

    // 创建一个新的 Excel 文件
    f, err := excelize.CreateFile()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create Excel file",
        })
        return
    }
    defer f.Close()

    // 设置 Excel 文件的标题
    f.SetSheetName(0, title)
    f.SetActiveSheet(0)

    // 向 Excel 文件添加数据
    // 这里只是一个示例，实际应用中应该根据需要添加数据
    f.SetCellValue("Sheet1", "A1", "Name")
    f.SetCellValue("Sheet1", "B1", "Age")
    for i := 2; i < 10; i++ {
        f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i), fmt.Sprintf("Person %d", i-1))
        f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i), 20+i)
    }

    // 生成 Excel 文件并下载
    if err := f.SaveAs("./output.xlsx"); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to save Excel file",
        })
        return
    }

    c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Header("Content-Disposition", "attachment; filename=output.xlsx")
    c.File("./output.xlsx")
}

func main() {
    r := gin.Default()
    // 使用 Gin 的 Logger 和 Recovery 中间件
    r.Use(gin.Logger(), gin.Recovery())

    // 定义一个路由，用于处理 Excel 文件生成的请求
    r.GET("/generate", ExcelGeneratorHandler)

    // 启动服务器
    r.Run() // 默认在 8080 端口
}
