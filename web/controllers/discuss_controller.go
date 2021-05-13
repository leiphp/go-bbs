package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取首页数据
func DiscussIndex(c *gin.Context){
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	c.HTML(http.StatusOK, "discuss/index.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
}

func DiscussList(c *gin.Context){
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	c.HTML(http.StatusOK, "discuss/list.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
}