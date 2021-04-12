package routers

import (
	"bbs/web/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.LoadHTMLGlob("/views/**/*")
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
		v1.GET("/topgoer", helloHandler)
		v1.GET("/test",  controllers.HomeIndex)
		v1.GET("/discuss/index",  controllers.DiscussIndex)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
	return r
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

