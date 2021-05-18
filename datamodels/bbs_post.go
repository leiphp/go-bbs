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
	CreateTime          int64   `json:"create_time"`           //创建时间
	CategoryId          int     `json:"category_id"`           //分类ID
	Remark              string  `json:"remark"`                //备注
	CategoryName        string  `gorm:"-" json:"category_name"`  //分类名称
	CommentCount        int     `gorm:"-" json:"comment_count"`  //评论量
	CreateDate          string  `gorm:"-" json:"create_date"`    //创建日期
	UserInfo		  struct{
		AuthorName        string  `json:"author_name"`           //作者昵称
		HeadImg           string  `json:"head_img"`              //作者头像
		IsVip     		  int     `json:"is_vip"`                //是否VIP
		IsAdmin     	  int     `json:"is_admin"`              //是否管理员
	} `gorm:"-" json:"user_info,omitempty"`                      //用户详情
}

//返回表名
func (this BbsPost) TableName() string {
	return "cy_post"
}

//组装数据，帖子类型
var PostType = map[int]string{
	1: "提问",
	2: "分享",
	3: "讨论",
	4: "建议",
	5: "公告",
	6: "动态",
}

//帖子评论列表参数
type ParamsPostCommentList struct {
	Page      	int64  `json:"page"`       //分页
	PerPage   	int64  `json:"per_page"`   //每页显示多少条
	PostId    	int    `json:"post_id"`    //文章ID
}

type BbsPostInfoVo struct {
	ID                int64   `json:"id"` 					 //ID
	Title          	  string  `json:"title"`            	 //标题
	Author            int     `json:"author"`                //作者ID
	Content           string  `json:"content"`               //内容
	Reward            int     `json:"reward"`                //奖励
	IsWonderful       int     `json:"is_wonderful"`          //是否精帖
	IsTop   	      int     `json:"is_top"`                //是否置顶
	Solved            int     `json:"solved"`                //是否解决
	ViewCount         int     `json:"view_count"`            //浏览量
	CommentCount      int     `json:"comment_count"`         //评论量
	CreateTime        int     `json:"create_time"`        //创建时间
	CreateDate        string  `json:"create_date"`        //创建日期
	CategoryId        int     `json:"category_id"`           //分类ID
	CategoryName      string  `json:"category_name"`         //分类名称
	UserInfo		  struct{
		AuthorName        string  `json:"author_name"`           //作者昵称
		HeadImg           string  `json:"head_img"`              //作者头像
		IsVip     		  int     `json:"is_vip"`                //是否VIP
		IsAdmin     	  int     `json:"is_admin"`              //是否管理员
	} `json:"user_info,omitempty"` //用户详情
}

//帖子列表参数
type PostPageListQuery struct {
	CategoryId    	*int   `json:"category_id"`   //分类id
	Page      	    int64  `json:"page"`          //分页
	PerPage   	    int64  `json:"per_page"`      //每页显示多少条
	Solved          *int   `json:"solved"`        //是否解决
	IsWonderful    	*int   `json:"is_wonderful"`  //是否精贴
}