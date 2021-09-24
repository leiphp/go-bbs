package repositories

import (
	"bbs/datamodels"
	"bbs/initialize"
	"github.com/go-redis/redis/v7"
	"strconv"
)


/*
	操作bbs_post表的接口定义

	作者名称：leixiaotian 创建时间：20210412
*/

type BbsPostInterface interface {
	SelectInfo(id int64) (datamodels.BbsPost, error) //获得帖子信息
	SelectTopList() ([]datamodels.BbsPost, error) //获得置顶帖子信息
	SelectPage(params map[string]interface{}, page int64, perPage int64) ([]datamodels.BbsPost, int, error) //获取分页帖子列表
	SelectByQuery(params map[string]interface{}, page int64, perPage int64, limit int64) ([]datamodels.BbsPost, int, error) //按条件获取帖子列表
}

//返回结构体对象
func NewBbsPost() BbsPostInterface {
	return &bbsPost{}
}

//bbsPost构体
type bbsPost struct {
}

//获得主播信息
func (this *bbsPost) SelectInfo(id int64) (datamodels.BbsPost, error) {

	var bbsPostInfo datamodels.BbsPost
	//redis礼物key
	//jmfMemberKey := ReturnRedisKey(API_CACHE_JMF_MEMBER, userId)
	//result, err := initialize.RedisCluster.Get(jmfMemberKey).Bytes()

	//读取数据库
	if err := initialize.MsqlDb.Where("id = ? ", id).Find(&bbsPostInfo).Error; err != nil {
		initialize.IrisLog.Errorf("[帖子仓库-获取帖子信息失败]-[%s]", err.Error())
		return bbsPostInfo, err
	}

	return bbsPostInfo, nil
}

//获得置顶帖子信息
func (this *bbsPost) SelectTopList() ([]datamodels.BbsPost, error) {
	type Top struct {
		Id  		int64 `json:"id"`
		CreateTime  int64 `json:"create_time"`
	}
	postTops := []Top{}
	var top Top
	var bbsPostInfo datamodels.BbsPost
	postList := []datamodels.BbsPost{}
	//redis礼物key
	giftKey := ReturnRedisKey(API_CACHE_POST_TOP_LIST, nil)
	initialize.IrisLog.Infof("[帖子仓库-获取redis置顶key]-[%s]", giftKey)
	list, err := initialize.RedisCluster.ZRevRange(giftKey, 0, 100).Result()
	initialize.IrisLog.Infof("[帖子仓库-获取redis置顶list]-[%s]", list)
	if err != nil {
		initialize.IrisLog.Errorf("[帖子仓库-获取redis帖子置顶列表失败]-[%s]", err.Error())
		return postList, err
	}

	//缓存数据为空
	if len(list) == 0 {
		//读取数据库
		if err := initialize.MsqlDb.Model(&bbsPostInfo).Where("is_top = ? ", 1).Find(&postList).Error; err != nil {
			initialize.IrisLog.Errorf("[帖子仓库-查询置顶帖子失败-%s]", err.Error())
			return postList, err
		}

		for _, val := range postList {
			top.Id = val.ID
			top.CreateTime = val.CreateTime
			postTops = append(postTops, top) //TODO 可省略
			initialize.RedisCluster.ZAdd(ReturnRedisKey(API_CACHE_POST_TOP_LIST, nil), &redis.Z{Score: float64(val.CreateTime), Member: val.ID})
		}

		//initialize.IrisLog.Infof("[帖子仓库-postTops]-[%s]", postTops)
		//设置redis缓存
		//jsonData, _ := json.Marshal(postTops)
		//initialize.IrisLog.Infof("[帖子仓库-jsonData]-[%s]", jsonData)
		//if err := initialize.RedisCluster.Set(ReturnRedisKey(API_CACHE_POST_TOP_LIST, nil), jsonData, 0).Err(); err != nil {
		//	initialize.IrisLog.Errorf("[帖子仓库-设置商城用户redis失败-%s]", err.Error())
		//	return postList, err
		//}
	}else {
		//val是礼物id
		for _, val := range list {
			postId, _ := strconv.ParseInt(val, 10, 64)
			postInfo, _ := this.SelectInfo(postId)
			//帖子加入列表
			postList = append(postList, postInfo)
		}
	}

	return postList, nil
}

//获取分页帖子列表
func (this *bbsPost) SelectPage(params map[string]interface{}, page int64, perPage int64) ([]datamodels.BbsPost, int, error) {

	var (
		info    datamodels.BbsPost
		records []datamodels.BbsPost
		total   = 0
	)

	db := initialize.MsqlDb.Model(&info).Select("cy_post.*, `cy_oauth_user`.nickname").Joins("left join cy_oauth_user ON cy_oauth_user.id = cy_post.author ")

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
		initialize.IrisLog.Errorf("[帖子仓库-获得帖子分页录列表-失败]-[%s]", err)
		return records, 0, err
	}
	return records, total, nil

}

//按条件获取帖子列表
func (this *bbsPost) SelectByQuery(params map[string]interface{}, page int64, perPage int64, limit int64) ([]datamodels.BbsPost, int, error) {

	var (
		info    datamodels.BbsPost
		records []datamodels.BbsPost
		total   = 0
	)

	db := initialize.MsqlDb.Model(&info).Select("cy_post.*, `cy_oauth_user`.nickname").Joins("left join cy_oauth_user ON cy_oauth_user.id = cy_post.author ")

	if params != nil {
		db = db.Where(params)
	}

	//总数
	db.Order("id DESC").Count(&total)

	if page > 0 && perPage > 0 {
		db = db.Limit(perPage).Offset((page - 1) * perPage)
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	err := db.Debug().Order("id DESC").Find(&records).Error
	if err != nil {
		initialize.IrisLog.Errorf("[帖子仓库-获得帖子分页录列表-失败]-[%s]", err)
		return records, 0, err
	}
	return records, total, nil

}