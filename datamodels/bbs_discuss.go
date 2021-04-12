package datamodels

/*
	讨论帖子表
	作者名称：leixiaotian 创建时间：20210412
*/
type BbsDiscuss struct {
	ID                  int     `gorm:"primary_key" json:"id"` //
	UserId              int64   `json:"user_id"`               //用户中心id
	GuildId             int64   `json:"guild_id"`              //工会id
	GuildName           string  `json:"guild_name"`            //工会名称
	GuildDivide         float64 `json:"guild_divide"`          //工会比例
	Divide              float64 `json:"divide"`                //分成比例
	TotalCommission     float64 `json:"total_commission"`      //历史佣金，总佣金
	CurrentCommission   float64 `json:"current_commission"`    //当前佣金
	DrawCommission      float64 `json:"draw_commission"`       //已提取佣金
	TotalDiamonds       float64 `json:"total_diamonds"`        //累计收益钻石
	CreateTime          int     `json:"create_time"`           //创建时间
	GuildAnchorDivide   float64 `json:"guild_anchor_divide"`   //工会比例
	TotalSellCommission float64 `json:"total_sell_commission"` //卖料佣金
	Drawings            int     `json:"drawings"`              //1可提现0不体现
	Freeze              int     `json:"freeze"`                //1禁播0可播
}

//返回表名
func (this BbsDiscuss) TableName() string {
	return "bbs_discuss"
}
