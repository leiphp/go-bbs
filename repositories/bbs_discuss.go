package repositories

import (
	"bbs/datamodels"

)


/*
	操作jmf_members表的接口定义

	作者名称：leixiaotian 创建时间：20210412
*/

type BbsDiscussInterface interface {
	SelectInfo(userId interface{}) (datamodels.BbsDiscuss, error) //获得主播信息
}

//返回结构体对象
func NewBbsDiscuss() BbsDiscussInterface {
	return &bbsDiscuss{}
}

//jmfMembers构体
type bbsDiscuss struct {
}

//获得主播信息
func (this *bbsDiscuss) SelectInfo(userId interface{}) (datamodels.BbsDiscuss, error) {

	var jmfMemberInfo datamodels.BbsDiscuss
	//redis礼物key
	/*jmfMemberKey := ReturnRedisKey(API_CACHE_JMF_MEMBER, userId)
	result, err := initialize.RedisCluster.Get(jmfMemberKey).Bytes()*/

	//读取数据库
	//if err := initialize.MsqlDb.Where("user_id = ? ", userId).Find(&jmfMemberInfo).Error; err != nil {
	//	initialize.IrisLog.Errorf("[获取主播信息失败]-[%s]", err.Error())
	//	return jmfMemberInfo, err
	//}


	return jmfMemberInfo, nil
}
