package repositories

import (
	"bbs/datamodels"
	"bbs/initialize"
)


/*
	操作bbs_user表的接口定义

	作者名称：leixiaotian 创建时间：20210513
*/

type BbsUserInterface interface {
	SelectInfo(id int64) (datamodels.BbsUser, error) //获得用户信息0
}

//返回结构体对象
func NewBbsUser() BbsUserInterface {
	return &bbsUser{}
}

//bbsUser构体
type bbsUser struct {
}

//获得用户信息
func (this *bbsUser) SelectInfo(id int64) (datamodels.BbsUser, error) {

	var userInfo datamodels.BbsUser
	//redis礼物key
	//jmfMemberKey := ReturnRedisKey(API_CACHE_JMF_MEMBER, userId)
	//result, err := initialize.RedisCluster.Get(jmfMemberKey).Bytes()

	//读取数据库
	if err := initialize.MsqlDb.Where("id = ? ", id).Find(&userInfo).Error; err != nil {
		initialize.IrisLog.Errorf("[获取主播信息失败]-[%s]", err.Error())
		return userInfo, err
	}

	return userInfo, nil
}
