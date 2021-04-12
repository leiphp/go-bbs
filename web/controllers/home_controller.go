package controllers

import (
	"bbs/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	首页控制器，提供首页相关操作

	作者名称：leixiaotian 创建时间：20210412
*/

//获取首页数据
func HomeIndex(c *gin.Context){
	result := make(map[int]string,0)
	result[0] = "leixiaotain"
	result[1] = "www.100txy.com"
	c.JSON(http.StatusOK, lib.ReturnJson(200, "", result))
}
