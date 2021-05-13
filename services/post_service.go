package services

import (
	"bbs/initialize"
	"bbs/libs"
	"bbs/repositories"
	"time"
)

/*
	提供关于讨论帖子服务

	作者名称：leixiaotian 创建时间：20210512
*/
type PostInterfaceService interface {
	GetPost(id int64) (interface{},error) //获取用户钱包
}

//初始化对象函数
func NewPostService() PostInterfaceService {
	return &postService{
		shopMemberService:       repositories.NewBbsDiscuss(),
		bbsPostService:          repositories.NewBbsPost(),
	}
}

type postService struct {
	shopMemberService 			repositories.BbsDiscussInterface     //商城会员服务
	bbsPostService 			    repositories.BbsPostInterface        //社区帖子服务
}

//获取用户钱包
func (this *postService) GetPost(id int64) (interface{},error){

	type BbsPostInfoVo struct {
		ID                int64   `json:"id"` 					 //ID
		Title          	  string  `json:"title"`            	 //标题
		Author            int     `json:"author"`                //作者ID
		Content           string  `json:"content"`               //内容
		Reward            int     `json:"reward"`                //奖励
		IsWonderful       int     `json:"is_wonderful"`          //是否精帖
		IsTop   	      int     `json:"is_top"`                //是否置顶
		Solved            int     `json:"solved"`                //是否解决
		ViewCount         int     `json:"view_count"`            //浏览量
		CommentCount      int     `json:"comment_count"`         //评论量
		CreateTime        int     `json:"create_time"`        //创建时间
		CreateDate        string  `json:"create_date"`        //创建日期
		CategoryId        int     `json:"category_id"`           //分类ID
		CategoryName      string  `json:"category_name"`         //分类名称
		UserInfo		  struct{
			AuthorName        string  `json:"author_name"`           //作者昵称
			HeadImg           string  `json:"head_img"`              //作者头像
			IsVip     		  int     `json:"is_vip"`                //是否VIP
			IsAdmin     	  int     `json:"is_admin"`              //是否管理员
		} `json:"user_info,omitempty"` //用户详情
	}
	var bbsPostInfoVo BbsPostInfoVo

	postInfo, err := this.bbsPostService.SelectInfo(id)
	initialize.IrisLog.Infof("[帖子服务-postInfo数据]-[%s]", libs.StructToJson(postInfo))
	if err != nil {
		return 3006, err
	}
	bbsPostInfoVo.ID = postInfo.ID
	bbsPostInfoVo.Title = postInfo.Title
	bbsPostInfoVo.Author = postInfo.Author
	bbsPostInfoVo.UserInfo.AuthorName = "雷小天"
	bbsPostInfoVo.UserInfo.HeadImg = "http://thirdqq.qlogo.cn/g?b=oidb&k=7iaib304zfK77M2ibtukgic1kQ&s=100&t=1585567893"
	bbsPostInfoVo.UserInfo.IsVip = 3
	bbsPostInfoVo.Content = postInfo.Content
	bbsPostInfoVo.Reward = postInfo.Reward
	bbsPostInfoVo.CreateDate = time.Unix(int64(postInfo.CreateTime), 0).Format("2006-01-02")
	bbsPostInfoVo.IsWonderful = postInfo.IsWonderful
	bbsPostInfoVo.IsTop = postInfo.IsTop
	bbsPostInfoVo.Solved = postInfo.Solved
	bbsPostInfoVo.ViewCount = postInfo.ViewCount
	bbsPostInfoVo.CommentCount = 12
	bbsPostInfoVo.CategoryName = "提问"

	return bbsPostInfoVo, nil
}
