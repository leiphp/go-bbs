package controllers

import (
	"bbs/datamodels"
	"bbs/initialize"
	"bbs/libs"
	"bbs/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostController struct {

	PostService services.PostInterfaceService
}

//获取帖子列表页
func (this *PostController) List(c *gin.Context){
	cate := c.Param("cate")
	page := c.DefaultQuery("page", "1")
	statusString := c.DefaultQuery("status", "0")
	status, _ := strconv.Atoi(statusString)
	fmt.Println("cate",cate)
	fmt.Println("page",page)
	category := map[string]int{
		"all":     0,  //全部
		"quiz":    1,  //提问
		"share":   2,  //分享
		"discuss": 3,  //讨论
		"suggest": 4,  //建议
		"notice":  5,  //公告
		"news":    6,  //动态
	}
	var params datamodels.PostPageListQuery
	params.Page,_ = strconv.ParseInt(page, 10, 64)
	params.PerPage = 20
	var cli int
	if cate != "all" {
		changeCate := category[cate]
		params.CategoryId = &changeCate
		cli = category[cate]
	}
	if status > 0 {
		switch status {
			case 1:
				val := 0
				params.Solved = &val
			case 2:
				val := 1
				params.Solved = &val
			case 3:
				val := 1
				params.IsWonderful = &val
		}
	}
	postList, _ := this.PostService.GetPostPageList(params)
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-获取postList数据]-[%s]", libs.StructToJson(postList))
	total := postList.(map[string]interface{})["total"] //map断言后再从map取值
	fmt.Println("total is:",total)
	pages := libs.CreatePaging(params.Page, params.PerPage, int64(total.(int)))
	fmt.Println("pages is:",pages)
	c.HTML(http.StatusOK, "post/list.html", gin.H{
		"title": "综合栏目-雷小天社区",
		"data": postList,
		"paging": pages,
		"cli": cli,
		"category": cate,
		"status": status,
		"prev": pages.Page - 1,
		"next": pages.Page + 1,
	})
	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"}))
	//c.HTML(http.StatusOK, "post/list.html", gin.H{"title": "社区讨论-雷小天社区", "address": "bbs.100txy.com"})
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
		"title": result.(datamodels.BbsPostInfoVo).Title,
		"address": "bbs.100txy.com",
		"id": id,
		"data": result,
		"comment": commentList,
	})
}