// 代码生成时间: 2025-08-04 08:02:22
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "sort"

    "github.com/gin-gonic/gin"
)

// FolderStructureOrganizer is a Gin handler that organizes a folder's structure.
func FolderStructureOrganizer(c *gin.Context) {
    // Get directory path from query parameter.
    dirPath := c.Query("path")
    if dirPath == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "path parameter is required",
        })
        return
    }

    // Check if the directory exists.
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": fmt.Sprintf("directory not found: %s", dirPath),
        })
        return
    }

    // Read the directory files and subdirectories.
    files, err := os.ReadDir(dirPath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("error reading directory: %s", err),
        })
        return
    }

    // Sort files and subdirectories by name.
    var fileList []os.DirEntry
    for _, file := range files {
        fileList = append(fileList, file)
    }
    sort.Slice(fileList, func(i, j int) bool {
        return fileList[i].Name() < fileList[j].Name()
    })

    // Prepare the response with a list of files and subdirectories.
    var response []string
    for _, file := range fileList {
        response = append(response, file.Name())
    }

    c.JSON(http.StatusOK, gin.H{
        "files": response,
    })
}

func main() {
    r := gin.Default()

    // Register the FolderStructureOrganizer handler with a path.
    r.GET("/organize", FolderStructureOrganizer)

    // Start the server.
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
