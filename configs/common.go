/*
@Time : 2021/04/12 14:45
@Author : LeiXiaoTian
@File : common.go
@Software: GoLand
*/
package configs

//定义公用常量
const (

	//一些常量
	Uniacid                       = 3    //挂接的公众号$uniacid
	ExchangeGoodsId               = 226  //兑换使用的商品ID
	CheckinGoodsId                = 227  //签到使用的商品ID
	RechargepackageAndroidGoodsId = 223  //安卓平台充值套餐使用的商品ID
	RechargepackageIosGoodsId     = 224  //ios平台充值套餐使用的商品ID
	RechargepackageWebGoodsId     = 225  //web平台充值套餐使用的商品ID
	VipPackageCategoryId          = 1214 //会员套餐使用的分类ID
	RechargePackageCategoryId     = 1215 //充值套餐使用的分类ID
	GiftsCategoryId               = 1217 //礼物分类Id
	DefaultPage                   = 1    //默认第一页
	DefaultPerPage                = 10   //默认一页10条记录

	//redisCacheKey
	CacheMemberInfo      = "MemberInfo:"     //用户详情
	CacheVipPackage      = "VipPackage"      //会员充值套餐
	CacheRechargePackage = "RechargePackage" //会员充值钻石套餐的配置信息
	CacheShopMember      = "ShopMember:"     //eweishopmember内容
	CacheGoods           = "Goods"           //商品详情
	//直播相关
	CacheJmfMemberDiamonds  = "Member:Diamond:" //单独维护钻石
	CacheJmfMember          = "JmfMemberInfo:"  //会员信息拓展
	CacheJmfGoodsGifts      = "Gift:Info:"      //礼物商品详情
	CacheJmfGoodsGiftsList  = "Gift:List:"      //礼物商品列表
	CacheJmfMemberCashsList = "Cash:List:"      //提现列表
	CacheJmfMemberCashsInfo = "Cash:Info:"      //提现明细

	CacheJmfMemberInfo  = "JmfMember"      //主播信息
	CacheJmfMemberUnion = "JmfMemberUnion" //工会信息

	//球豆相关
	BeansOpen       = "Setting:BallGold:WithDraw" //是否开启球豆提现
	WebDiamondsList = "DiamondRecord:Unified:"    //专门给web使用钻石消费，充值明细

	Page    = 1
	PerPage = 10

)

//定义消息容器
var MsgCode map[int]string = map[int]string{
	200: "成功",
	400: "客户端请求的语法错误，服务器无法理解",
	403: "服务器理解请求客户端的请求，但是拒绝执行此请求",
	404: "请求的资源不存在",
	405: "客户端请求中的方法被禁止",
	408: "服务器等待客户端发送的请求时间过长，超时",
	500: "内部服务器错误",
	501: "服务器不支持请求的功能，无法完成请求",

	//参数相关
	400100: "请检查请求参数",
	400101: "",
	//用户相关
	400200: "用户不存在",
	400201: "自己不用送礼物给自己",
	400202: "送礼物用户不存在",
	400203: "接受礼物用户不存在",
	400204: "送礼用户钻石不足",
	400205: "主播不存在",
	400206: "修改主播分成比例失败",
	400207: "修改工会比例失败",
	400208: "获取会员列表失败",
	400209: "请登录在操作",
	400210: "非法操作，你的操作已被记录，请勿继续尝试",
	400211: "获取主播收益列表失败",
	400212: "获得主播指定时间钻石数量失败",
	400213: "获得排行榜数据失败",

	//商品相关
	400300: "礼物不存在",
	400301: "礼物已下架",
	400302: "获得礼物列表失败",
	400303: "礼物名称只能6个字符",
	400304: "后台添加礼物失败",
	400305: "钻石余额不足不能送礼物",
	//银行相关
	400400: "银行支行不存在",
	400401: "绑定银行卡失败",
	400402: "您已经绑定过银行卡",
	400403: "该银行卡已经被绑定过，请使用其他卡号",
	400404: "用户银行卡信息不存在",
	//提现相关
	400500: "提现失败",
	400501: "提取的佣金大于余额",
	400502: "审核失败",
	400503: "提现时间应该在当月20-25号",
	400504: "提现金额不满足条件",
	400506: "抱歉，您目前的角色不能提现",
	400507: "提现余额不足",
	400508: "提现金额异常",
	400509: "该用户还未绑定支付宝，不能提现",
	400510: "当天提现次数已达三次",
	400511: "该主播设置为不能提现",
	400512: "工会主播不能提现",
	400513: "该提现类型未授权绑定",

	//评论相关
	400600: "查询评论列表失败",

	//系统相关
	500100: "系统错误",
}
