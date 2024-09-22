package main

import (
    "university-distribution-backend/config"
    "university-distribution-backend/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors" // 导入 CORS 包
)

func main() {
    // 初始化数据库连接
    config.InitDB()

    // 创建 Gin 引擎
    r := gin.Default()

    // 设置静态文件路径（assets 文件夹）
    r.Static("/assets", "./assets")

    // 配置 CORS 中间件
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // 允许的来源
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // 注册路由
    routes.SetupRoutes(r)

    // 启动服务器
    r.Run(":8080")
}
