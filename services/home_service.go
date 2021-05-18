package services

import (
	"bbs/datamodels"
	"bbs/repositories"
	"math"
	"time"
)

/*
	提供关于首页相关服务

	作者名称：leixiaotian 创建时间：20210513
*/
type HomeInterfaceService interface {
	GetTopList() (interface{},error) //获取置顶帖子
	GetPost(id int64) (interface{},error) //获取帖子详情
	GetPostList(query datamodels.PostPageListQuery) (interface{}, error) //获取所有分页帖子
}

//初始化对象函数
func NewHomeService() HomeInterfaceService {
	return &homeService{
		bbsUserService:          repositories.NewBbsUser(),
		shopMemberService:       repositories.NewBbsDiscuss(),
		bbsPostService:          repositories.NewBbsPost(),
	}
}

type homeService struct {
	bbsUserService 			    repositories.BbsUserInterface         //社区会员服务
	shopMemberService 			repositories.BbsDiscussInterface     //商城会员服务
	bbsPostService 			    repositories.BbsPostInterface        //社区帖子服务
}

//获取置顶帖子
func (this *homeService) GetTopList() (interface{},error){
	var bbsPostInfoVo datamodels.BbsPostInfoVo
	var bbsPostInfoVoList []datamodels.BbsPostInfoVo
	list, err := this.bbsPostService.SelectTopList()
	if err != nil {
		return 8004, err
	}
	for _, val := range list {
		userInfo,_ :=this.bbsUserService.SelectInfo(int64(val.Author))
		bbsPostInfoVo.ID = val.ID
		bbsPostInfoVo.Title = val.Title
		bbsPostInfoVo.CategoryName = datamodels.PostType[val.CategoryId]
		bbsPostInfoVo.Solved = val.Solved
		bbsPostInfoVo.Reward = val.Reward
		bbsPostInfoVo.CreateDate = time.Unix(int64(val.CreateTime), 0).Format("2006-01-02")
		bbsPostInfoVo.CommentCount = 12
		bbsPostInfoVo.UserInfo.HeadImg = userInfo.HeadImg
		bbsPostInfoVo.UserInfo.AuthorName = userInfo.Nickname
		bbsPostInfoVo.UserInfo.IsVip = userInfo.IsVip
		bbsPostInfoVo.UserInfo.IsAdmin = userInfo.IsAdmin
		bbsPostInfoVoList = append(bbsPostInfoVoList,bbsPostInfoVo)
	}
	return bbsPostInfoVoList, nil
}


//获取帖子详情
func (this *homeService) GetPost(id int64) (interface{},error){

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
		CreateTime        int64   `json:"create_time"`           //创建时间
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
	bbsPostInfoVo.CategoryName = datamodels.PostType[postInfo.CategoryId]

	return bbsPostInfoVo, nil
}

// 获取所有分页帖子
func (this *homeService) GetPostList(query datamodels.PostPageListQuery) (interface{}, error) {

	//如果没有分页，默认是第一页和显示20条
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PerPage == 0 {
		query.PerPage = 20
	}

	params := make(map[string]interface{})
	if query.CategoryId != nil {
		params["category_id"] = *query.CategoryId
	}
	if query.Solved != nil {
		params["solved"] = *query.Solved
	}
	if query.IsWonderful != nil {
		params["is_wonderful"] = *query.IsWonderful
	}

	list, total, err := this.bbsPostService.SelectPage(params, query.Page, query.PerPage)
	if err != nil {
		return 5001, err
	}

	for key, val := range list {
		userInfo,_ :=this.bbsUserService.SelectInfo(int64(val.Author))
		list[key].CategoryName = datamodels.PostType[val.CategoryId]
		list[key].CommentCount = 12
		list[key].CreateDate = time.Unix(int64(val.CreateTime), 0).Format("2006-01-02")
		list[key].UserInfo.HeadImg = userInfo.HeadImg
		list[key].UserInfo.AuthorName = userInfo.Nickname
		list[key].UserInfo.IsVip = userInfo.IsVip
		list[key].UserInfo.IsAdmin = userInfo.IsAdmin
	}

	//分页返回
	result := map[string]interface{}{
		"count":   len(list),                                             //当前页面多少条
		"total":   total,                                                 //记录总数
		"pages":   math.Round(float64(total)/float64(query.PerPage)) + 1, //总共多少页
		"page":    query.Page,                                            //当前页数
		"perPage": query.PerPage,                                         //每页多少条
		"rows":    list,
	}

	return result, err
}
