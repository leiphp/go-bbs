package main

import (
	"bbs/routers"
	"fmt"
)

func main() {
	// 初始化路由
	r := routers.SetupRouter()

	//加载模板文件目录
	r.LoadHTMLGlob("views/**/*")

	err := r.Run(":8001")
	if  err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
