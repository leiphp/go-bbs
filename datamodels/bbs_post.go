package datamodels

/*
	讨论帖子表
	作者名称：leixiaotian 创建时间：20210512
*/
type BbsPost struct {
	ID                  int64   `gorm:"primary_key" json:"id"` //ID
	Title          		string  `json:"title"`            	   //标题
	Author              int     `json:"author"`                //作者ID
	Content             string  `json:"content"`               //内容
	Reward              int     `json:"reward"`                //奖励
	IsWonderful         int     `json:"is_wonderful"`          //是否精帖
	IsVip     		    int     `json:"is_vip"`                //是否VIP
	IsTop   			int     `json:"is_top"`                //是否置顶
	Solved              int     `json:"solved"`                //是否解决
	ViewCount           int     `json:"view_count"`            //浏览量
	CreateTime          int     `json:"create_time"`           //创建时间
	CategoryId          int     `json:"category_id"`           //分类ID
	Remark              string  `json:"remark"`                //备注
}

//返回表名
func (this BbsPost) TableName() string {
	return "cy_post"
}
