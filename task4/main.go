package main

import (
	"github.com/gin-gonic/gin"
	"go_learn/task4/database"
	"go_learn/task4/routes"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 创建Gin实例
	router := gin.Default()

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务
	router.Run(":8080")
}
