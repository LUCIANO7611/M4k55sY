// 代码生成时间: 2025-09-20 02:24:59
package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// RenameRequest 定义了重命名请求的数据结构
type RenameRequest struct {
    SourcePath string   `json:"source_path"`
    NewNames   []string `json:"new_names"`
}

func main() {
    r := gin.Default()
    
    // 注册批量重命名的处理器
    r.POST("/rename", func(c *gin.Context) {
        var req RenameRequest
        
        // 绑定JSON请求体到RenameRequest结构体
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request format",
            })
            return
        }
        
        // 检查源路径是否存在
        if _, err := os.Stat(req.SourcePath); os.IsNotExist(err) {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Source path does not exist",
            })
            return
        }
        
        // 检查新名称的数量是否与文件数量匹配
        files, err := os.ReadDir(req.SourcePath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to read source path",
            })
            return
        }
        if len(files) != len(req.NewNames) {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "The number of new names does not match the number of files",
            })
            return
        }
        
        // 进行文件重命名操作
        for i, file := range files {
            if file.IsDir() {
                continue // 跳过目录
            }
            oldPath := filepath.Join(req.SourcePath, file.Name())
            newPath := filepath.Join(req.SourcePath, req.NewNames[i])
            if err := os.Rename(oldPath, newPath); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": fmt.Sprintf("Failed to rename %s to %s: %v", oldPath, newPath, err),
                })
                return
            }
        }
        
        // 如果所有文件都重命名成功，返回成功响应
        c.JSON(http.StatusOK, gin.H{
            "message": "Files renamed successfully",
        })
    })
    
    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        fmt.Printf("Failed to start server: %v
", err)
    }
}
