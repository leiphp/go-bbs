package routers

import (
	"bbs/services"
	"bbs/web/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	homeController 	*controllers.HomeController
	postController 	*controllers.PostController
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.LoadHTMLGlob("/views/**/*")
	// 路由组1 ，处理GET请求
	v1 := r.Group("/")

	//初始化控制器结构体
	initControllerStruct()
	//quizService := services.NewPostService()
	//obj := controllers.PostController{
	//	PostService: quizService,
	//}
	// {} 是书写规范
	{
		v1.GET("/",  homeController.Index)
		v1.GET("/home/list",  homeController.List)
		v1.GET("/discuss",  controllers.DiscussIndex)
		v1.GET("/discuss/list",  controllers.DiscussList)
		v1.GET("/post/list/:cate",  postController.List)
		v1.GET("/post/:id",  postController.Detail)
		v1.GET("/login", login)
		v1.GET("submit", submit)
		v1.GET("/topgoer", helloHandler)
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	v3 := r.Group("/check")
	{
		v3.GET("/health", func(c *gin.Context) {
			//c.String(http.StatusOK, "hello word")
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": []int{},
				"msg": "service is ok",
			})
		})
	}

	//r.Run(":8000")
	return r
}

//初始化所以控制器结构体
func initControllerStruct() {
	//首页控制器
	homeController = homeObj()
	//帖子控制器
	postController = postObj()
}

//	首页控制器结构体
func homeObj() *controllers.HomeController {
	homeService := services.NewHomeService()
	obj := controllers.HomeController{
		HomeService: homeService,
	}
	return &obj
}

//	帖子控制器结构体
func postObj() *controllers.PostController {
	postService := services.NewPostService()
	obj := controllers.PostController{
		PostService: postService,
	}
	return &obj
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

