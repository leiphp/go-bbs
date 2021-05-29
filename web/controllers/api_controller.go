package controllers

import (
	"bbs/libs"
	"bbs/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct {

	ApiService services.ApiInterfaceService
}

func NewApiController() *ApiController {
	obj := &ApiController{
		ApiService: services.NewApiService(),
	}
	return obj
}

//获取帖动态消息
func (this *ApiController) MessageData(c *gin.Context){
	type message struct {
		Avatar    	  string		`json:"avatar"`
		Content   	  string		`json:"content"`
		NowUserLink   string		`json:"now_user_link"`
		Type 		  string		`json:"type"`
	}

	message1 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://thirdqq.qlogo.cn/g?b=oidb&k=FV064fMPPEKRbPeiaEJMJdw&s=100&t=1556912115\" class=\"avatar avatar-22123 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"leo_221\"/>",
		Content: "<a href=\"https://q.jinsom.cn/author/22123\" target=\"_blank\">11加入了<n>LightSNS官网</n></a>",
		NowUserLink: "https://q.jinsom.cn/author/22123",
		Type: "reg"}
	message2 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://img.jinsom.cn/user_files/16386/avatar/29761083_1594000978.jpg\" class=\"avatar avatar-16386 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"参天大树\"/>",
		Content: "<a href=\"https://q.jinsom.cn/38123.html\" target=\"_blank\">22回复了<n><font style=\"color:#FF5722;\" class=\"vip-user user-1\">jinsom</font></n>的帖子</a>",
		NowUserLink: "https://q.jinsom.cn/author/16386",
		Type: "comment-bbs"}
	message3 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://thirdqq.qlogo.cn/g?b=oidb&k=WKiaCOEuVWRPvzpDPUOBSQQ&s=100&t=1556314539\" class=\"avatar avatar-21729 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"༺ཌ滕ད༻\"/>",
		Content: "<a href=\"https://q.jinsom.cn/32519.html\" target=\"_blank\">33购买了付费内容</a>",
		NowUserLink: "https://q.jinsom.cn/author/21729",
		Type: "buy"}
	message4 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://thirdqq.qlogo.cn/g?b=oidb&k=WKiaCOEuVWRPvzpDPUOBSQQ&s=100&t=1556314539\" class=\"avatar avatar-21729 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"༺ཌ滕ད༻\"/>",
		Content: "<a href=\"https://q.jinsom.cn/32519.html\" target=\"_blank\">44购买了付费内容</a>",
		NowUserLink: "https://q.jinsom.cn/author/21729",
		Type: "buy"}
	message5 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://thirdqq.qlogo.cn/g?b=oidb&k=WKiaCOEuVWRPvzpDPUOBSQQ&s=100&t=1556314539\" class=\"avatar avatar-21729 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"༺ཌ滕ད༻\"/>",
		Content: "<a href=\"https://q.jinsom.cn/32519.html\" target=\"_blank\">55购买了付费内容</a>",
		NowUserLink: "https://q.jinsom.cn/author/21729",
		Type: "buy"}
	message6 := &message{
		Avatar: "<img loading=\"lazy\" src=\"https://thirdqq.qlogo.cn/g?b=oidb&k=WKiaCOEuVWRPvzpDPUOBSQQ&s=100&t=1556314539\" class=\"avatar avatar-21729 avatar-normal opacity\" width=\"40\" height=\"40\" alt=\"༺ཌ滕ད༻\"/>",
		Content: "<a href=\"https://q.jinsom.cn/32519.html\" target=\"_blank\">66购买了付费内容</a>",
		NowUserLink: "https://q.jinsom.cn/author/21729",
		Type: "buy"}
	messages := []*message{message1, message2, message3,message4, message5, message6}

	//测试推送MQ消息
	this.ApiService.PushVisitLog(1)

	//result := make(map[int]string,0)
	//result[0] = "leixiaotain"
	//result[1] = "www.100txy.com"
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", messages))
}
