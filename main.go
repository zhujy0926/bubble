package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASES bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
