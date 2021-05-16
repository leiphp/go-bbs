package controllers

import (
	"bbs/datamodels"
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
func (this *PostController) List(c *gin.Context){
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	c.HTML(http.StatusOK, "post/list.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
}

//获取帖子详情页
func (this *PostController) Detail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.IrisLog.Infof("[帖子控制器-PostDetail-http请求数据]-[%d]", id)

	result, err := this.PostService.GetPost(int64(id))
	initialize.IrisLog.Infof("[帖子控制器-PostDetail-post返回数据]-[%s]", libs.StructToJson(result))
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404.html", gin.H{"title": "404"})
		return
	}
	//获取评论
	var query datamodels.ParamsPostCommentList
	query.PostId = id
	commentList, _ := this.PostService.GetPostCommentList(query)
	initialize.IrisLog.Infof("[帖子控制器-commentList返回数据]-[%s]", libs.StructToJson(commentList))
	c.HTML(http.StatusOK, "post/detail.html", gin.H{
		"title": "社区讨论-雷小天社区",
		"address": "bbs.100txy.com",
		"id": id,
		"data": result,
		"comment": commentList,
	})
}