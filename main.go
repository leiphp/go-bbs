package main

import (
	"bbs/routers"
	"fmt"
)

func main() {
	// 1.创建路由
	r := routers.SetupRouter()

	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	//// 1.创建路由
	//r := gin.Default()
	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	//r.Run(":8000")
}
