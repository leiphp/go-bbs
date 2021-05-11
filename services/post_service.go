package services

import (
	"bbs/repositories"
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
		AuthorName        string  `json:"author_name"`           //作者昵称
		Content           string  `json:"content"`               //内容
		Reward            int     `json:"reward"`                //奖励
		IsWonderful       int     `json:"is_wonderful"`          //是否精帖
		IsVip     		  int     `json:"is_vip"`                //是否VIP
		IsTop   	      int     `json:"is_top"`                //是否置顶
		Solved            int     `json:"solved"`                //是否解决
		ViewCount         int     `json:"view_count"`            //浏览量
		CreateTime        int     `json:"create_time"`           //创建时间
		CategoryId        int     `json:"category_id"`           //分类ID
		CategoryName      string  `json:"category_name"`         //分类名称
	}
	var bbsPostInfoVo BbsPostInfoVo

	postInfo, err := this.bbsPostService.SelectInfo(id)
	if err != nil {
		return 3006, err
	}
	bbsPostInfoVo.ID = postInfo.ID
	bbsPostInfoVo.Title = postInfo.Title
	bbsPostInfoVo.Author = postInfo.Author
	bbsPostInfoVo.AuthorName = "雷小天"
	bbsPostInfoVo.Content = postInfo.Content
	bbsPostInfoVo.Reward = postInfo.Reward
	bbsPostInfoVo.CreateTime = postInfo.CreateTime
	bbsPostInfoVo.IsWonderful = postInfo.IsWonderful
	bbsPostInfoVo.IsTop = postInfo.IsTop
	bbsPostInfoVo.Solved = postInfo.Solved
	bbsPostInfoVo.ViewCount = postInfo.ViewCount
	bbsPostInfoVo.CategoryName = "提问"

	return bbsPostInfoVo, nil
}
