package controllers

import (
	"bbs/datamodels"
	"bbs/initialize"
	"bbs/libs"
	"bbs/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	首页控制器，提供首页相关操作

	作者名称：leixiaotian 创建时间：20210412
*/

type HomeController struct {

	HomeService services.HomeInterfaceService
}

//获取首页数据
func (this *HomeController) Index(c *gin.Context){
	type Params struct {
		ID                int     `gorm:"primary_key" json:"id"` //
		UserId            int64   `json:"user_id"`               //用户中心id
		ExpertId          int64   `json:"expert_id"`             //专家id
		TotalCommission   float64 `json:"total_commission"`      //历史佣金，总佣金
		CurrentCommission float64 `json:"current_commission"`    //当前佣金
	}
	var query1 Params
	query22 := &Params{ID:222}
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-http请求数据]-[%s]", libs.StructToJson(query1))
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-http请求数据]-[%s]", libs.StructToJson(query22))

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

	type reply struct {
		Id		   int
		HeadImg    string
		Nickname   string
		ReplyNum   int
	}

	type comment struct {
		Id		   int
		Title    string
		CommentNum   int
	}

	//置顶数据
	topData, err := this.HomeService.GetTopList()
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-获取置顶帖子数据]-[%s]", libs.StructToJson(topData))
	if err != nil {
		c.HTML(http.StatusNotFound, "error/404.html", gin.H{"title": "404"})
	}

	//分页列表
	var postQuery datamodels.PostPageListQuery
	postList, err := this.HomeService.GetPostList(postQuery)
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-获取params]-[%s]", libs.StructToJson(postQuery))
	initialize.IrisLog.Infof("[主页控制器-HomeIndex-获取postList数据]-[%s]", libs.StructToJson(postList))

	reply1 := &reply{Id: 1, HeadImg: "http://thirdqq.qlogo.cn/g?b=oidb&k=7iaib304zfK77M2ibtukgic1kQ&s=100&t=1585567893", Nickname: "雷小天", ReplyNum: 32}
	reply2 := &reply{Id: 2, HeadImg: "http://tvax2.sinaimg.cn/crop.28.0.300.300.50/90af5a89ly8fi45eieuewj20b408cdg2.jpg", Nickname: "雷小天_", ReplyNum: 12}
	reply3 := &reply{Id: 3, HeadImg: "http://thirdqq.qlogo.cn/qqapp/101370818/676C95317C3A2B80E5CE2C532894C83C/100", Nickname: "雷小天科技", ReplyNum: 10}
	reply4 := &reply{Id: 4, HeadImg: "http://thirdqq.qlogo.cn/g?b=oidb&k=7iaib304zfK77M2ibtukgic1kQ&s=100&t=1585567893", Nickname: "雷小天博客", ReplyNum: 9}
	reply5 := &reply{Id: 5, HeadImg: "https://bbs.100txy.com/public/static/res/images/avatar/default.png", Nickname: "霸王不别姬", ReplyNum: 6}
	reply6 := &reply{Id: 6, HeadImg: "http://tva1.sinaimg.cn/crop.12.0.75.75.180/ac64bb8fjw8e7wm6xwlq7j202s023a9v.jpg", Nickname: "梦星空A", ReplyNum: 5}
	reply7 := &reply{Id: 7, HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/1918095D33E92FBF50BB993FFD9A2BA8/100", Nickname: "灿若繁星", ReplyNum: 4}
	reply8 := &reply{Id: 8, HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/7FF8706F8DDBEB51F548C4C6CB28509B/100", Nickname: "Aries", ReplyNum: 3}
	reply9 := &reply{Id: 9, HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/007645B246880D226C409199E9420F66/100", Nickname: "一米阳光", ReplyNum: 2}
	reply10 := &reply{Id: 10, HeadImg: "https://tva1.sinaimg.cn/crop.0.0.118.118.180/5db11ff4gw1e77d3nqrv8j203b03cweg.jpg", Nickname: "贤心", ReplyNum: 1}
	reply11 := &reply{Id: 11, HeadImg: "http://qzapp.qlogo.cn/qzapp/101370818/D770E84703CDF381F35C49D660A6CC39/100", Nickname: "呵呵", ReplyNum: 1}
	reply12 := &reply{Id: 12, HeadImg: "http://thirdqq.qlogo.cn/g?b=oidb&k=7iaib304zfK77M2ibtukgic1kQ&s=100&t=1585567893", Nickname: "雷小天2", ReplyNum: 1}

	comment1 := &comment{Id: 1, Title: "什么是同步和异步，阻塞和非阻塞", CommentNum: 32}
	comment2 := &comment{Id: 2, Title: " linux下开启、关闭、重启mysql服务", CommentNum: 12}
	comment3 := &comment{Id: 3, Title: " Linux关闭防火墙命令", CommentNum: 10}
	comment4 := &comment{Id: 4, Title: " MySQL主从复制的原理解析", CommentNum: 9}
	comment5 := &comment{Id: 5,Title: "PHP连接MySQL数据库的三种方式(mysql、mysqli、pdo)", CommentNum: 6}
	comment6 := &comment{Id: 4, Title: " 微信小程序文字超出限制如何在末尾加省略号", CommentNum: 2}

	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "首页-雷小天社区",
		"topdata": topData,
		"data": postList,
		"reply": []*reply{reply1, reply2, reply3, reply4, reply5, reply6, reply7, reply8, reply9, reply10, reply11, reply12},
		"comment": []*comment{comment1, comment2, comment3, comment4, comment5, comment6},
	})
}

func (this *HomeController) Reg(c *gin.Context){
	c.HTML(http.StatusOK, "home/reg.html", gin.H{"title": "首页-雷小天社区", "address": "www.100txy.com"})
}

func (this *HomeController) List(c *gin.Context){
	c.HTML(http.StatusOK, "home/list.html", gin.H{"title": "首页-雷小天社区", "address": "www.100txy.com"})
}

func (this *HomeController) Cate(c *gin.Context){
	result := make(map[int]string,0)
	result[0] = "leixiaotain"
	result[1] = "www.100txy.com"
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", result))
}