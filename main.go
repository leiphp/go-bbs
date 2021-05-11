package main

import (
	"bbs/configs"
	"bbs/initialize"
	"bbs/routers"
	"fmt"
)

func main() {
	//获得配置对象
	Yaml := configs.InitConfig()
	initialize.Init(Yaml)

	//注册nacos
	go func() {
		initialize.InitRegisterServer()
	}()

	// 初始化路由
	r := routers.SetupRouter()

	//加载静态文件夹
	r.Static("/assets", "./static")

	//加载模板文件目录
	r.LoadHTMLGlob("views/**/*")

	err := r.Run(":8001")
	if  err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
