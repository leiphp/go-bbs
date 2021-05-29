package services

import (
	"bbs/datamodels"
	"bbs/initialize"
	"bbs/libs"
	"bbs/repositories"
	"errors"
	"math"
	"time"
)

/*
	提供关于API服务

	作者名称：leixiaotian 创建时间：20210529
*/
type ApiInterfaceService interface {
	GetPost(id int64) (interface{},error) //获取论坛帖子详情
	GetPostCommentList(query datamodels.ParamsPostCommentList) (interface{}, error)    //获得帖子评论
	GetPostPageList(query datamodels.PostPageListQuery) (interface{}, error) //获取所有分页帖子
}

//初始化对象函数
func NewApiService() ApiInterfaceService {
	return &apiService{
		bbsUserService:          repositories.NewBbsUser(),
		bbsPostService:          repositories.NewBbsPost(),
		bbsCommentService:       repositories.NewBbsComment(),
	}
}

type apiService struct {
	bbsUserService 			    repositories.BbsUserInterface           //社区会员服务
	bbsPostService 			    repositories.BbsPostInterface           //社区帖子服务
	bbsCommentService 			repositories.BbsCommentInterface        //社区评论服务
}

//获取用户钱包
func (this *apiService) GetPost(id int64) (interface{},error){
	var bbsPostInfoVo datamodels.BbsPostInfoVo

	postInfo, err := this.bbsPostService.SelectInfo(id)
	initialize.IrisLog.Infof("[帖子服务-postInfo数据]-[%s]", libs.StructToJson(postInfo))
	if err != nil {
		initialize.IrisLog.Errorf("[帖子服务-获取帖子信息失败]-[%s]", err.Error())
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

// 获得帖子评论
func (this *apiService) GetPostCommentList(query datamodels.ParamsPostCommentList) (interface{}, error) {

	//如果没有分页，默认是第一页和显示20条
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PerPage == 0 {
		query.PerPage = 20
	}

	params := make(map[string]interface{})
	params["postid"] = query.PostId

	list, total, err := this.bbsCommentService.SelectPage(params, query.Page, query.PerPage)
	if err != nil {
		return 400600, errors.New("论坛评论服务-获取帖子评论列表失败！")
	}

	for key,val := range list {
		userInfo,_ :=this.bbsUserService.SelectInfo(int64(val.Ouid))
		list[key].CreateDate = time.Unix(val.CreateTime,0).Format("2006-01-02")
		list[key].HeadImg = userInfo.HeadImg
		list[key].Nickname = userInfo.Nickname
		list[key].IsAdmin = userInfo.IsAdmin
		list[key].IsVip = userInfo.IsVip
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

// 获取所有分页帖子
func (this *apiService) GetPostPageList(query datamodels.PostPageListQuery) (interface{}, error) {

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