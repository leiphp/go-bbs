package datamodels

/*
	讨论评论表
	作者名称：leixiaotian 创建时间：20210512
*/
type BbsComment struct {
	ID                  int64   `gorm:"primary_key" json:"id"` //ID
	Ouid          		int     `json:"ouid"`            	   //作者ID
	Like                int     `json:"like"`                  //喜欢
	Pid                 int     `json:"pid"`                   //父级评论id
	Postid              int     `json:"postid"`                //帖子id
	Content             string  `json:"content"`               //评论内容
	CreateTime          int64   `json:"create_time"`           //创建时间
}

//返回表名
func (this BbsComment) TableName() string {
	return "cy_comment"
}
