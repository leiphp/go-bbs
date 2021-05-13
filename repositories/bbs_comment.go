package repositories

import (
	"bbs/datamodels"
	"bbs/initialize"
)


/*
	操作bbs_comment表的接口定义

	作者名称：leixiaotian 创建时间：20210513
*/

type BbsCommentInterface interface {
	SelectPage(params map[string]interface{}, page int64, perPage int64) ([]datamodels.BbsComment, int, error) //查询帖子评论列表
}

//返回结构体对象
func NewBbsComment() BbsCommentInterface {
	return &bbsComment{}
}

//bbsComment构体
type bbsComment struct {
}

//查询帖子评论列表
func (this *bbsComment) SelectPage(params map[string]interface{}, page int64, perPage int64) ([]datamodels.BbsComment, int, error) {

	var (
		info    datamodels.BbsComment
		records []datamodels.BbsComment
		total   = 0
	)

	db := initialize.MsqlDb.Model(&info)

	if params != nil {
		db = db.Where(params)
	}
	//总数
	db.Order("id DESC").Count(&total)

	if page > 0 && perPage > 0 {
		db = db.Limit(perPage).Offset((page - 1) * perPage)
	}

	err := db.Order("id DESC").Find(&records).Error
	if err != nil {
		initialize.IrisLog.Errorf("[帖子评论仓库-获得评论列表失败]-[%s]", err)
		return records, 0, err
	}
	return records, total, nil

}
