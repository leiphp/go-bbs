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

type UserController struct {

	UserService services.UserInterfaceService
}

func NewUserController() *UserController {
	obj := &UserController{
		UserService: services.NewUserService(),
	}
	return obj
}


//获取用户详情页
func (this *UserController) Detail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.IrisLog.Infof("[用户控制器-http请求数据]-[%d]", id)

	result, err := this.UserService.GetUserInfo(int64(id))
	initialize.IrisLog.Infof("[用户控制器-userInfo返回数据]-[%s]", libs.StructToJson(result))
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404.html", gin.H{"title": "404"})
		return
	}

	//获取发布的帖子
	var query datamodels.PostByQuery
	uid := int64(id)
	query.UserId = &uid
	query.Limit = 15
	postList, _ := this.UserService.GetPostByQuery(query)
	initialize.IrisLog.Infof("[用户控制器-postList返回数据]-[%s]", postList)

	//获取发布的评论
	//var query datamodels.ParamsPostCommentList
	//query.PostId = id
	//commentList, _ := this.PostService.GetPostCommentList(query)
	//initialize.IrisLog.Infof("[帖子控制器-commentList返回数据]-[%s]", libs.StructToJson(commentList))

	c.HTML(http.StatusOK, "user/detail.html", gin.H{
		"title": result.(datamodels.BbsUser).Nickname,
		"address": "bbs.100txy.com",
		"id": id,
		"data": result,
		"postData": postList,
	})
}