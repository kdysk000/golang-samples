package main

import (
	"gorm-sample/controller"
	"gorm-sample/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	mysql.DbInit()
	mysql.AutoMigrate()

	engine := gin.Default()
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.POST("/add", controller.UserAdd)
			v1.GET("/list", controller.UserList)
			v1.PUT("/update", controller.UserUpdate)
			v1.DELETE("/delete", controller.UserDelete)
		}
	}
	engine.Run(":3000")
}
