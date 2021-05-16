package repositories

import (
	"bbs/initialize"
	"fmt"
	"github.com/go-redis/redis/v7"
	"math"
)

/*
	redis缓存的key拼接

	作者名称：LeiXiaoTian 创建时间：20210516
*/

const KeyPrefix = "GoCommunity:"

const (
	API_CACHE_POST_TOP_LIST             = 1  //帖子置顶列表
	API_CACHE_VIP_PACKAGE               = 2  //会员充值套餐
	API_CACHE_RECHARGE_PACKAGE          = 3  //会员充值钻石套餐的配置信息
	API_CACHE_EXCHANGE_PACKAGE          = 4  //钻石兑换金币配置信息
	API_CACHE_CHECKIN_CONFIG            = 5  //会员签到配置信息
	API_CACHE_BALLGOLD_HISTORY          = 6  //球豆变更列表
	API_CACHE_BALLGOLD                  = 7  //球豆明细
	API_CACHE_EXCHANGE_HISTORY          = 8  //钻石兑换列表
	API_CACHE_EXCHANGE                  = 9  //钻石兑换明细
	API_CACHE_CHECKIN_HISTORY           = 10 //签到列表
	API_CACHE_CHECKIN                   = 11 //签到明细
	API_CACHE_RECHARGE_HISTORY          = 12 //钻石充值列表
	API_CACHE_RECHARGE                  = 13 //钻石充值明细
	API_CACHE_MEMBER                    = 14 //mc_members内容
	API_CACHE_SHOP_MEMBER               = 15 //ewei_shop_member内容
	API_CACHE_TASK_COMPLETE             = 16 //任务完成记录
	API_CACHE_GOODS                     = 17 //商品详情
	API_CACHE_GOODSOPTION               = 18 //商品的某个具体的选项
	API_CACHE_GOODSPARAMS               = 19 //商品的参数列表
	API_CACHE_ORDER                     = 20 //订单详情
	API_CACHE_ORDER_TRANSID             = 21 //Transaction id 与order no之间的对应关系
	API_CACHE_ORDER_GOODS               = 22 //订单的商品详情（第一个商品）
	API_CACHE_VIP_HISTORY               = 23 //购买套餐列表
	API_CACHE_GOLD_HISTORY              = 24 //金币变更列表
	API_CACHE_DIAMOND_HISTORY           = 25 //钻石变更列表
	API_CACHE_CREDIT_RECORD             = 26 //钻石、金币、球豆明细（使用的同一个表存储的)
	API_CACHE_ALIPAY_QR_URL             = 27 //支付宝二维码支付链接
	API_CACHE_CONSUME_HISTORY           = 28 //钻石消费列表
	API_CACHE_GIFT_LIST                 = 29 //礼物列表
	API_CACHE_GIFT_INFO                 = 30 //礼物详情
	API_CACHE_MEMBER_DIAMOND            = 31 //用户钻石数，单独维护
	API_CACHE_MEMBER_LOCK               = 32 //用户操作的互斥锁，比如防止同一个用户创建多次
	API_CACHE_SETTING_BALLGOLD_WITHDRAW = 33 //球豆是否可以提现的设置
	API_CACHE_IOS_WAITPAY_ORDERS        = 34 //苹果应用内支付代付款的订单列表
	API_CACHE_DIAMOND_HISTORY_UNIFIED   = 35 //统一的钻石记录明细，包括充值和消费
	API_CACHE_GIFT_RECORD_SENDER        = 36 //礼物详情的送礼物部分

)

var apiCacheKeys = map[int]string{
	API_CACHE_POST_TOP_LIST:             "Post:Top:List",             //用户详情
	API_CACHE_VIP_PACKAGE:               "VipPackage",                //会员充值套餐
	API_CACHE_RECHARGE_PACKAGE:          "RechargePackage",           //会员充值钻石套餐的配置信息
	API_CACHE_EXCHANGE_PACKAGE:          "ExchangePackage",           //钻石兑换金币配置信息
	API_CACHE_CHECKIN_CONFIG:            "CheckInConfig",             //会员签到配置信息
	API_CACHE_BALLGOLD_HISTORY:          "BallGoldHistory",           //球豆变更列表
	API_CACHE_BALLGOLD:                  "BallGold",                  //球豆明细
	API_CACHE_EXCHANGE_HISTORY:          "ExchangeHistory",           //钻石兑换列表
	API_CACHE_EXCHANGE:                  "Exchange",                  //钻石兑换明细
	API_CACHE_CHECKIN_HISTORY:           "CheckinHistory",            //签到列表
	API_CACHE_CHECKIN:                   "CheckIn",                   //签到明细
	API_CACHE_RECHARGE_HISTORY:          "RechargeHistory",           //钻石充值列表
	API_CACHE_CONSUME_HISTORY:           "ConsumeHistory",            //钻石消费列表
	API_CACHE_RECHARGE:                  "Recharge",                  //钻石充值明细
	API_CACHE_MEMBER:                    "Member",                    //mc_members内容
	API_CACHE_SHOP_MEMBER:               "ShopMember",                //ewei_shop_member内容
	API_CACHE_TASK_COMPLETE:             "TaskComplete",              //任务完成记录
	API_CACHE_GOODS:                     "Goods",                     //商品详情
	API_CACHE_GOODSOPTION:               "GoodsOption",               //商品的某个具体的选项
	API_CACHE_GOODSPARAMS:               "GoodsParams",               //商品的某个具体的参数
	API_CACHE_ORDER:                     "Order",                     //订单详情
	API_CACHE_ORDER_TRANSID:             "Order:Trans",               //Transaction id 与order no之间的对应关系
	API_CACHE_ORDER_GOODS:               "Order:Goods",               //订单的商品详情（第一个商品）
	API_CACHE_VIP_HISTORY:               "VIPHistory",                //购买套餐列表
	API_CACHE_GOLD_HISTORY:              "GoldHistory",               //金币变更列表
	API_CACHE_DIAMOND_HISTORY:           "DiamondHistory",            //钻石变更列表
	API_CACHE_CREDIT_RECORD:             "CreditRecord",              //钻石、金币、球豆明细（使用的同一个表存储的)
	API_CACHE_ALIPAY_QR_URL:             "Alipay:QRURL",              //支付宝二维码支付链接
	API_CACHE_GIFT_LIST:                 "Gift:List",                 //礼物列表
	API_CACHE_GIFT_INFO:                 "Gift:Info",                 //礼物详情
	API_CACHE_MEMBER_DIAMOND:            "Member:Diamond",            //用户钻石数，单独维护
	API_CACHE_MEMBER_LOCK:               "Member:Lock",               //用户操作的互斥锁，比如防止同一个用户创建多次
	API_CACHE_SETTING_BALLGOLD_WITHDRAW: "Setting:BallGold:WithDraw", //球豆是否可以提现的设置
	API_CACHE_IOS_WAITPAY_ORDERS:        "Orders:IOS:WaitPay",        //苹果应用内支付代付款的订单列表
	API_CACHE_DIAMOND_HISTORY_UNIFIED:   "DiamondRecord:Unified",     //统一的钻石记录明细，包括充值和消费
	API_CACHE_GIFT_RECORD_SENDER:        "Gift:Record",               //礼物详情的送礼物部分

}

//分页结构
type RedisPage struct {
	Count   int      `json:"count"`    //当前页面多少条
	Total   int64    `json:"total"`    //记录总数
	Pages   float64  `json:"pages"`    //总共多少页
	Page    int64    `json:"page"`     ////当前页数
	PerPage int64    `json:"per_page"` //每页多少条
	Rows    []string `json:"rows"`     //每页多少条
}

/**
 * 获取对应的key
 */
func ReturnRedisKey(keyType int, key interface{}) string {

	var suffix string
	if key != nil {
		suffix = ":" + fmt.Sprintf("%v", key)
	}
	redisKey := KeyPrefix + apiCacheKeys[keyType] + suffix
	return redisKey
}

/**
GetRedisPage function
redis分页封装
*/
func GetRedisPage(redisKey string, page int64, perPage int64, min string, max string) (RedisPage, error) {

	var pageList RedisPage

	if min == "" && max == "" {
		min = "0"          //sorce的值最小0
		max = "9999999999" //sorce得值最大9千万
	}

	limitStart := (page - 1) * perPage
	limitEnd := (limitStart + perPage) - 1

	rangeBy := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: limitStart,
		Count:  limitEnd,
	}

	rows, err := initialize.RedisCluster.ZRevRangeByScore(redisKey, &rangeBy).Result()
	if err != nil {
		initialize.IrisLog.Errorf("[redis分页数据出错]-[%s]", err.Error())
		return pageList, err
	}
	//统计ScoreSet集合元素总数
	count := initialize.RedisCluster.ZCount(redisKey, min, max).Val()

	pageList.Count = len(rows)
	pageList.Total = count
	pageList.Pages = math.Round(float64(count)/float64(perPage)) + 1
	pageList.Page = page
	pageList.PerPage = perPage
	pageList.Rows = rows

	return pageList, nil
}
