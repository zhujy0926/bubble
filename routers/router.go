package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		// 代办事项

		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id")
	}
	return r
}
