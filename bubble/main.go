package main

import (
	"fmt"
	"hellp/dao"
	"hellp/models"
	"hellp/routers"
)

func main() {
	// gorm 框架
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)

	}
	fmt.Println("连接成功")
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// gin 框架
	r := routers.SetupRouter()
	r.Run(":8080")

}
