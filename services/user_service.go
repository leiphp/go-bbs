package services

import (
	"bbs/datamodels"
	"bbs/initialize"
	"bbs/libs"
	"bbs/repositories"
	"math"
	"time"
)

/*
	提供关于用户会员服务

	作者名称：leixiaotian 创建时间：20210530
*/
type UserInterfaceService interface {
	GetUserInfo(id int64) (interface{},error) //获取用户详情
	GetPostByQuery(query datamodels.PostByQuery) (interface{}, error) //获取满足条件帖子
	GetPostPageList(query datamodels.PostPageListQuery) (interface{}, error) //获取所有分页用户
}

//初始化对象函数
func NewUserService() UserInterfaceService {
	return &userService{
		bbsUserService:          repositories.NewBbsUser(),
		bbsPostService:          repositories.NewBbsPost(),
		bbsCommentService:       repositories.NewBbsComment(),
	}
}

type userService struct {
	bbsUserService 			    repositories.BbsUserInterface           //社区会员服务
	bbsPostService 			    repositories.BbsPostInterface           //社区帖子服务
	bbsCommentService 			repositories.BbsCommentInterface        //社区评论服务
}

//获取用户钱包
func (this *userService) GetUserInfo(id int64) (interface{},error){

	userInfo, err := this.bbsUserService.SelectInfo(id)
	initialize.IrisLog.Infof("[用户服务-userInfo数据]-[%s]", libs.StructToJson(userInfo))
	if err != nil {
		initialize.IrisLog.Errorf("[用户服务-获取用户信息失败]-[%s]", err.Error())
		return 3006, err
	}
	userInfo.CreateDate = time.Unix(userInfo.CreateTime, 0).Format("2006-01-02 15:04:05")
	return userInfo, nil
}

// 获取所有分页用户
func (this *userService) GetPostByQuery(query datamodels.PostByQuery) (interface{}, error) {

	//如果没有分页，默认是第一页和显示20条
	//if query.Page == 0 {
	//	query.Page = 1
	//}
	//if query.PerPage == 0 {
	//	query.PerPage = 20
	//}

	params := make(map[string]interface{})
	if query.UserId != nil {
		params["author"] = *query.UserId
	}
	if query.CategoryId != nil {
		params["category_id"] = *query.CategoryId
	}
	if query.Solved != nil {
		params["solved"] = *query.Solved
	}
	if query.IsWonderful != nil {
		params["is_wonderful"] = *query.IsWonderful
	}

	list, total, err := this.bbsPostService.SelectByQuery(params, query.Page, query.PerPage, query.Limit)
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

// 获取所有分页用户
func (this *userService) GetPostPageList(query datamodels.PostPageListQuery) (interface{}, error) {

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