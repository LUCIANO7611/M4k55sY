// 代码生成时间: 2025-08-20 10:26:58
package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    \_ "github.com/go-sql-driver/mysql" // 导入MySQL驱动
)

// DB 是数据库连接的全局变量
var DB *sql.DB

// 函数 initDB 初始化数据库连接
func initDB() error {
    const (
        host     = "your_host"
        port     = 3306
        user     = "your_username"
        password = "your_password"
        dbname   = "your_dbname"
    )
    // 注意：需替换成实际的数据库连接信息
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(10)
    DB.SetConnMaxLifetime(time.Hour)
    return nil
}

// GinWithSQLInjectionProtect 是防止SQL注入的Gin处理器
func GinWithSQLInjectionProtect(c *gin.Context) {
    var query string
    // 这里假设前端传递了一个名为 "search" 的查询参数
    search := c.Param("search")
    // 使用strings.Fields函数来分割字符串，防止SQL注入
    search = strings.Join(strings.Fields(search), " \")
    query = `SELECT * FROM users WHERE username LIKE ?`
    // 使用参数化查询来防止SQL注入
    rows, err := DB.Query(query, "%"+search+"%")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal Server Error",
        })
        return
    }
    defer rows.Close()
    // 处理查询结果...
    c.JSON(http.StatusOK, gin.H{
        "message": "Query executed successfully",
    })
}

func main() {
    err := initDB()
    if err != nil {
        fmt.Println("Database connection failed: ", err)
        return
    }
    router := gin.Default()
    // 使用Gin中间件来记录日志
    router.Use(gin.Logger())
    // 使用Gin中间件来恢复任何panic
    router.Use(gin.Recovery())
    // 设置防止SQL注入的路由
    router.GET("/users/search/:search", GinWithSQLInjectionProtect)
    // 启动Gin服务器
    router.Run(":8080")
}
