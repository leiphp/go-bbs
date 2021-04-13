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
	type post struct {
		Id		   int
		Author     string
		HeadImg    string
		Nickname   string
		Name       string
		Title      string
		IsAdmin    int8
		IsVip      int8
		CreatTime  int64
		CreatDate  string
		Reward     int
		Solved     int8
		CommentNum int

	}
	post1 := &post{Id: 1, Author: "雷小天", HeadImg: "http://thirdqq.qlogo.cn/g?b=oidb&k=7iaib304zfK77M2ibtukgic1kQ&s=100&t=1585567893",
		Nickname: "雷小天", Name: "提问", Title: "Vue 项目中如何去除url中的 #", IsAdmin: 0, IsVip:3, CreatTime: 1618328126, CreatDate: "2020-01-19", Reward: 20, Solved: 1, CommentNum: 2}
	post2 := &post{Id: 2, Author: "呵呵", HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/D770E84703CDF381F35C49D660A6CC39/100",
		Nickname: "呵呵", Name: "讨论", Title: "组装电脑，主板用哪个的兼容好点，不考虑什么超频", IsAdmin: 0, IsVip:0, CreatTime: 1618328126, CreatDate: "2018-03-02", Reward: 20, Solved: 1, CommentNum: 1}
	post3 := &post{Id: 3, Author: "Aries", HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/7FF8706F8DDBEB51F548C4C6CB28509B/100",
		Nickname: "Aries", Name: "提问", Title: "什么是同步和异步，阻塞和非阻塞", IsAdmin: 0, IsVip:3, CreatTime: 1618328126, CreatDate: "2018-03-02", Reward: 20, Solved: 1, CommentNum: 3}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "首页-雷小天社区",
		"topdata": [3]*post{post1, post2, post3},
	})
}

func HomeList(c *gin.Context){
	c.HTML(http.StatusOK, "home_list.html", gin.H{"title": "首页-雷小天社区", "address": "www.100txy.com"})
}

func HomeCate(c *gin.Context){
	result := make(map[int]string,0)
	result[0] = "leixiaotain"
	result[1] = "www.100txy.com"
	c.JSON(http.StatusOK, lib.ReturnJson(200, "", result))
}