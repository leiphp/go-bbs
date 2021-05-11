package controllers

import (
	"bbs/initialize"
	"bbs/libs"
	"bbs/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController struct {

	PostService services.PostInterfaceService
}

//获取帖子列表页
func (this *PostController) PostList(c *gin.Context){
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	c.HTML(http.StatusOK, "post_list.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
}

//获取帖子详情页
func (this *PostController) PostDetail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.IrisLog.Infof("[帖子控制器-PostDetail-http请求数据]-[%d]", id)

	result, err := this.PostService.GetPost(int64(id))
	initialize.IrisLog.Infof("[帖子控制器-PostDetail-post返回数据]-[%s]", libs.StructToJson(result))
	if err != nil {
		c.HTML(http.StatusNotFound, "user/404.html", gin.H{"title": "404"})
	}
	c.HTML(http.StatusOK, "post_detail.html", gin.H{
		"title": "社区讨论-雷小天社区",
		"address": "bbs.100txy.com",
		"id": id,
		"info": result,
	})
}