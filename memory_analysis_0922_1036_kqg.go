// 代码生成时间: 2025-09-22 10:36:36
 * and return them in a JSON format.
 */

package main

import (
    "fmt"
    "net/http"
    "runtime"
    "github.com/gin-gonic/gin"
)

// MemoryUsageResponse defines the structure for the memory usage response.
type MemoryUsageResponse struct {
    Alloc     uint64 `json:"alloc"`     // Bytes allocated and not yet freed.
    TotalAlloc uint64 `json:"totalAlloc"` // Total bytes allocated during the execution.
    Sys       uint64 `json:"sys"`       // Total bytes obtained from the OS.
    Malcs     uint64 `json:"malcs"`     // Number of mallocs.
    Freems    uint64 `json:"freems"`    // Number of frees.
    LiveObjs  uint64 `json:"liveObjs"`  // Number of live objects.
    Pauses    []uint64 `json:"pauses"`    // Pause times for GC.
    NumGC     uint32 `json:"numGC"`     // Number of GC runs.
}

// getMemoryUsage retrieves the current memory usage statistics.
func getMemoryUsage() MemoryUsageResponse {
    m := &runtime.MemStats{}
    runtime.ReadMemStats(m)
    return MemoryUsageResponse{
# 改进用户体验
        Alloc:     m.Alloc,
        TotalAlloc: m.TotalAlloc,
        Sys:       m.Sys,
        Malcs:     m.Mallocs,
        Freems:    m.Frees,
        LiveObjs:  m.HeapObjects,
        Pauses:    m.PauseNs[:],
        NumGC:     m.NumGC,
    }
}

// MemoryUsageHandler handles requests for memory usage analysis.
// It reads the current memory statistics and returns them in a JSON response.
func MemoryUsageHandler(c *gin.Context) {
    memoryUsage := getMemoryUsage()
# FIXME: 处理边界情况
    if err := c.JSON(http.StatusOK, memoryUsage); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
# 添加错误处理
            "error": fmt.Sprintf("Failed to return memory usage data: %v", err),
        })
    }
}

func main() {
    r := gin.Default()
    r.GET("/memory", MemoryUsageHandler)
    // You can add more handlers or middlewares here if needed.
# 添加错误处理
    r.Run() // listen and serve on 0.0.0.0:8080
}
