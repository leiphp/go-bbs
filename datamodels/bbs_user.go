package datamodels

/*
	用户表
	作者名称：leixiaotian 创建时间：20210512
*/
type BbsUser struct {
	ID                  int64   `gorm:"primary_key" json:"id"` //ID
	Uid          		int     `json:"uid"`            	   //UID
	Type                int     `json:"type"`                  //类型
	Nickname            string  `json:"nickname"` 	           //昵称
	HeadImg             string  `json:"head_img"` 	           //用户头像
	Openid              string  `json:"openid"`                //openid
	AccessToken         string  `json:"access_token"`          //access_token
	CreateTime          int64   `json:"create_time"`           //创建时间
	LastLoginTime       int64   `json:"last_login_time"`       //最后登录时间
	LastLoginIp         string  `json:"last_login_ip"`         //最后登录ip
	LoginTimes          int     `json:"login_times"` 	       //登录次数
	Status              int     `json:"status"` 	           //状态
	Email               string  `json:"email"`                 //邮箱
	IsAdmin             int     `json:"is_admin"` 	           //是否管理员
	IsVip               int     `json:"is_vip"` 	           //是否会员
	Reward              int64   `json:"reward"` 	           //奖励
	Pass               string  `json:"pass"`                 //密码
}

//返回表名
func (this BbsUser) TableName() string {
	return "cy_oauth_user"
}
