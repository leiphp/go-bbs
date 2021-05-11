package controllers

import (
	"bbs/initialize"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//获取帖子列表页
func PostList(c *gin.Context){
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	c.HTML(http.StatusOK, "post_list.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
}

//获取帖子详情页
func PostDetail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.IrisLog.Infof("[帖子控制器-PostDetail-http请求数据]-[%d]", id)

	c.HTML(http.StatusOK, "post_detail.html", gin.H{
		"title": "社区讨论-雷小天社区",
		"address": "bbs.100txy.com",
		"id": id,
	})
}