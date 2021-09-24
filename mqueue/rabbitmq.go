package mqueue

import (
	"bbs/datamodels"
	"bbs/initialize"
	"bbs/repositories"
	"encoding/json"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

type RabbitInterfaceService interface {
	PushVisitLog(userInfo datamodels.BbsUser)         //推送浏览日志
	//pushUserExp(params datamodels.SendGiftParams, giftInfo datamodels.Gift)           //推送用户等级
	//pushShopOrder(params datamodels.SendGiftParams, giftInfo datamodels.Gift)         //推送商城订单
	//pushLiveGift(params datamodels.SendGiftParams, giftInfo datamodels.Gift)          //推送直播礼物
	//PullShopOrder(number int)                                                         //处理消费订单
}

//初始化对象函数
func NewRabbitService() RabbitInterfaceService {
	return &rabbitService{
		userService:                  repositories.NewBbsUser(),
		//giftService:                  repositories.NewGift(),
		//memberService:                repositories.NewMember(),
		//shopMemberService:            repositories.NewShopMember(),
		//jmfMemberService:             repositories.NewJmfMembers(),
		//jmfMemberUnionService:        repositories.NewJmfMembersUnion(),
		//jmfLiveExchangeService:       repositories.NewJmfLiveExchange(),
		//jmfLiveExchangeLogService:    repositories.NewJmfLiveExchangeLog(),
		//memberBackpackService:        repositories.NewMembersBackpack(),
		//jmfLiveExchangeParamsService: repositories.NewJmfLiveExchangeParams(),
		//shopGiftService:              repositories.NewShopGift(),
		//jmfMembersBackpackService:    repositories.NewMembersBackpack(),
	}
}

type rabbitService struct {
	userService                  repositories.BbsUserInterface
	//giftService                  repositories.GiftInterface
	//memberService                repositories.MembersInterface
	//shopMemberService            repositories.ShopMemberInterface
	//jmfMemberService             repositories.JmfMembersInterface
	//jmfMemberUnionService        repositories.JmfMembersUnionInterface
	//jmfLiveExchangeService       repositories.JmfLiveExchangeInterface
	//jmfLiveExchangeLogService    repositories.JmfLiveExchangeLogInterface
	//memberBackpackService        repositories.MembersBackpackInterface
	//jmfLiveExchangeParamsService repositories.JmfLiveExchangeParamsInterface
	//shopGiftService              repositories.ShopGiftInterface
	//jmfMembersBackpackService    repositories.MembersBackpackInterface
}

//推送浏览日志
func (this *rabbitService) PushVisitLog(userInfo datamodels.BbsUser) {
	//新建channel
	visitLogCh, err := initialize.MqClientUCenter.GetChannel()
	if err != nil {
		initialize.IrisLog.Errorf("[推送浏览日志-创建channel失败]-[%s]", err.Error())
		return
	}
	defer visitLogCh.Close()

	//绑定队列
	q, queueErr := visitLogCh.QueueDeclare(
		initialize.Config.GetString("RabbitMQ.visitLogQueue"), //name
		true,  //durable
		false, //delete when usused
		false, //exclusive
		false, //nowait
		nil,   //argments
	)
	if queueErr != nil {
		initialize.IrisLog.Errorf("[推送浏览日志-创建queue失败]-[%s]", err.Error())
		return
	}

	//用户经验结构体
	type ExpInfo struct {
		FromUid       int64       `json:"fromUid"`       //发送用户id
		ToUid         int64       `json:"toUid"`         //接收用户id
		SourceType    int64       `json:"sourceType"`    //1充值消费2是任务
		EventType     string      `json:"eventType"`     //事件类型: live_send_diamond_gift|直播间赠送蓝钻礼物
		UnitType      int         `json:"unitType"`      //换算单位: 1|次数 2|蓝钻 3|球豆
		UnitNum       float64     `json:"unitNum"`       //单位数量
		InnerUserFlag bool        `json:"innerUserFlag"` //是否公司员工
		BizType       int         `json:"bizType"`       //业务类型
		BizId         string      `json:"bizId"`         //业务id
		RoomId        string      `json:"roomId"`        //房间id
		Date          string      `json:"date"`          //日期
		OtherInfo     interface{} `json:"otherInfo,omitempty"`
	}

	//封装body
	var exp ExpInfo
	var isStaff bool
	exp.FromUid = userInfo.ID
	exp.ToUid = 2
	exp.SourceType = 1
	//蓝钻礼物
	exp.EventType = "live_send_diamond_gift"
	exp.UnitType = 2
	exp.UnitNum = 20
	isStaff = true
	exp.InnerUserFlag = isStaff
	exp.BizId = strconv.Itoa(120)
	exp.BizType = 1
	exp.RoomId = strconv.FormatInt(110, 10)
	exp.Date = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	//封装body

	jsonBytes, err := json.Marshal(exp)
	if err != nil {
		initialize.IrisLog.Errorf("[推送浏览日志-json格式化失败]-[%s]", err.Error())
		return
	}

	//推送信息
	pushErr := visitLogCh.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        jsonBytes,
		})

	if pushErr != nil {
		initialize.IrisLog.Errorf("[推送浏览日志-推送失败]-[%s]", err.Error())
		return
	}

	initialize.IrisLog.Infof("[推送浏览日志-推送成功]-[%s]", string(jsonBytes))
}

////推送用户等级
//func (this *rabbitService) pushUserExp(params datamodels.SendGiftParams, giftInfo datamodels.Gift) {
//	//新建channel
//	userExpCh, err := initialize.MqClientUCenter.GetChannel()
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送用户等级-创建channel失败]-[%s]", err.Error())
//		return
//	}
//	defer userExpCh.Close()
//
//	//绑定队列
//	q, queueErr := userExpCh.QueueDeclare(
//		initialize.Config.GetString("RabbitMQ.UserExpQueue"), //name
//		true,  //durable
//		false, //delete when usused
//		false, //exclusive
//		false, //nowait
//		nil,   //argments
//	)
//	if queueErr != nil {
//		initialize.IrisLog.Errorf("[推送用户等级-创建queue失败]-[%s]", err.Error())
//		return
//	}
//
//	//用户经验结构体
//	type ExpInfo struct {
//		FromUid       int64       `json:"fromUid"`       //发送用户id
//		ToUid         int64       `json:"toUid"`         //接收用户id
//		SourceType    int64       `json:"sourceType"`    //1充值消费2是任务
//		EventType     string      `json:"eventType"`     //事件类型: live_send_diamond_gift|直播间赠送蓝钻礼物
//		UnitType      int         `json:"unitType"`      //换算单位: 1|次数 2|蓝钻 3|球豆
//		UnitNum       float64     `json:"unitNum"`       //单位数量
//		InnerUserFlag bool        `json:"innerUserFlag"` //是否公司员工
//		BizType       int         `json:"bizType"`       //业务类型
//		BizId         string      `json:"bizId"`         //业务id
//		RoomId        string      `json:"roomId"`        //房间id
//		OtherInfo     interface{} `json:"otherInfo,omitempty"`
//	}
//
//	//发送礼物用户的信息
//	sendMember, err := this.shopMemberService.SelectInfo(params.SendId)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送用户等级-查询发送礼物用户信息失败]-[%s]", err.Error())
//		return
//	}
//
//	//封装body
//	var exp ExpInfo
//	var isStaff bool
//	exp.FromUid = params.SendId
//	exp.ToUid = params.ReceiveId
//	exp.SourceType = 1
//	//蓝钻礼物
//	exp.EventType = "live_send_diamond_gift"
//	exp.UnitType = 2
//	exp.UnitNum = params.GiftCount * giftInfo.Diamonds
//	if sendMember.IsStaff == 1 {
//		isStaff = true
//	}
//	exp.InnerUserFlag = isStaff
//	exp.BizId = strconv.Itoa(params.GiftId)
//	exp.BizType = 1
//	exp.RoomId = strconv.FormatInt(params.RoomId, 10)
//	//封装body
//
//	jsonBytes, err := json.Marshal(exp)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送用户等级-json格式化失败]-[%s]", err.Error())
//		return
//	}
//
//	//推送信息
//	pushErr := userExpCh.Publish(
//		"",     // exchange
//		q.Name, // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "text/json",
//			Body:        jsonBytes,
//		})
//
//	if pushErr != nil {
//		initialize.IrisLog.Errorf("[推送用户等级-推送失败]-[%s]", err.Error())
//		return
//	}
//
//	initialize.IrisLog.Infof("[推送用户等级-推送成功]-[%s]", string(jsonBytes))
//}
//
////推送商城订单
//func (this *rabbitService) pushShopOrder(params datamodels.SendGiftParams, giftInfo datamodels.Gift) {
//
//	shopOrderCh, err := initialize.MqClient.GetChannel()
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-创建channel失败]-[%s]", err.Error())
//		return
//	}
//	defer shopOrderCh.Close()
//
//	q, err := shopOrderCh.QueueDeclare(
//		initialize.Config.GetString("RabbitMQ.SendGiftQueue"), //name
//		true,  //durable
//		false, //delete when usused
//		false, //exclusive
//		false, //nowait
//		nil,   //argments
//	)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-创建queue失败]-[%s]", err.Error())
//		return
//	}
//
//	jsonBytes, err := json.Marshal(params)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-json格式化失败]-[%s]", err.Error())
//		return
//	}
//
//	//发送礼物用户的信息
//	_, sendErr := this.shopMemberService.SelectInfo(params.SendId)
//	if sendErr != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-查询发送礼物用户信息失败]-[%s]", sendErr.Error())
//		return
//	}
//
//	err = shopOrderCh.Publish(
//		"",     // exchange
//		q.Name, // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "text/json",
//			Body:        jsonBytes,
//		})
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-推送失败]-[%s]", err.Error())
//		return
//	}
//
//	//发送参数入库
//	info := datamodels.JmfLiveExchangeParams{}
//	info.ID = 0
//	info.SendId = params.SendId
//	info.GoodsId = params.GiftId
//	info.ReceiveId = params.ReceiveId
//	info.SendTime = params.SendTime
//	info.RoomId = params.RoomId
//	info.Remark = string(jsonBytes)
//	if err := this.jmfLiveExchangeParamsService.Insert(info); err != nil {
//		initialize.IrisLog.Errorf("[推送商城订单-请求参数插入数据库失败]-[%s]", err.Error())
//	}
//
//	initialize.IrisLog.Infof("[推送商城订单-推送成功]-[%s]", string(jsonBytes))
//}
//
////推送直播礼物
//func (this *rabbitService) pushLiveGift(params datamodels.SendGiftParams, giftInfo datamodels.Gift) {
//
//	liveCh, err := initialize.MqClient.GetChannel()
//	if err != nil {
//		initialize.IrisLog.Infof("[推送粉丝经验-创建channel失败]-[%s]", err.Error())
//		return
//	}
//	defer liveCh.Close()
//
//	//主播信息
//	anchorsMember, err := this.jmfMemberService.SelectInfo(params.ReceiveId)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-获取主播信息失败]-[%s]", err.Error())
//		return
//	}
//
//	//发送礼物用户的信息
//	sendMember, err := this.shopMemberService.SelectInfo(params.SendId)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-查询发送礼物用户信息失败]-[%s]", err.Error())
//		return
//	}
//
//	//接收礼物用户的信息
//	receiveMember, err := this.shopMemberService.SelectInfo(params.ReceiveId)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-查询接收礼物用户信息失败]-[%s]", err.Error())
//		return
//	}
//
//	//主播是否含有工会信息
//	var memberUnion datamodels.JmfMembersUnion
//	if anchorsMember.GuildId > 0 {
//		memberUnion, err = this.jmfMemberUnionService.SelectInfo(anchorsMember.GuildId)
//		if err != nil {
//			initialize.IrisLog.Errorf("[推送粉丝经验-查询工会信息失败]-[%s]", err.Error())
//			return
//		}
//	}
//
//	//给直播平台返回信息
//	type LiveGift struct {
//		SendUid            int64   `json:"send_uid"`             //发送礼物用户
//		SendNickName       string  `json:"send_nickname"`        //发送礼物用户昵称
//		SendUserImg        string  `json:"send_user_img"`        //发送礼物用户头像
//		ReceiveUid         int64   `json:"receive_uid"`          //接收礼物用户
//		ReceiveNickName    string  `json:"receive_nickname"`     //接收礼物用户昵称
//		ReceiveUserImg     string  `json:"receive_user_img"`     //接收礼物用户头像
//		GiftId             int     `json:"gift_id"`              //礼物id
//		GiftName           string  `json:"gift_name"`            //礼物名称
//		GiftPrice          float64 `json:"gift_price"`           //蓝钻礼物价格
//		BallGoldPrice      float64 `json:"gift_ballgold_price"`  //球豆礼物价格
//		GiftType           string  `json:"gift_price_type_name"` //礼物价格类型
//		GiftCount          float64 `json:"gift_count"`           //礼物数量
//		RoomId             int64   `json:"room_id"`              //房间id
//		ReceiveCount       float64 `json:"receive_count"`        //扣款成功接收者收到多少
//		GiftIcon           string  `json:"gift_icon"`            //礼物图标
//		GiftEffect         string  `json:"gift_effect"`          //礼物动
//		GiftEffectImage    string  `json:"gift_effect_image"`    //礼物动效图
//		EffectWebp         string  `json:"gift_effect_webp"`     //webp礼物动效图
//		GiftPlayTime       int     `json:"gift_play_time"`       //webp礼物播放时长
//		GiftTypeId         int     `json:"gift_type_id"`         //礼物类型名称
//		SendUserLevel      int     `json:"send_user_level"`      //用户等级 默认不用
//		SendUserType       int     `json:"send_user_type"`       //用户等级
//		ReceiveUserLevel   int     `json:"receive_user_level"`   //用户等级 默认不用
//		ReceiveUserType    int     `json:"receive_user_type"`    //用户等级
//		ExchangeCommission float64 `json:"number_commission"`    //此个礼物主播的收益
//		GuildDivide        float64 `json:"guild_divide"`         //工会比例
//		Divide             float64 `json:"divide"`               //平台比例
//		IsStaff            int     `json:"is_staff"`             //是否公司员工
//		GiftBackpackId     int64   `json:"gift_backpack_id"`     //礼物背包id
//		GiftHotValue       float64 `json:"gift_hot_value"`       //礼物热度值
//		GiftCategory       int     `json:"gift_category"`        //礼物类型
//		GiftSendType       int     `json:"gift_send_type"`
//		SceneType     	   int     `json:"scene_type"`           //场景id【1：单聊】
//	}
//
//	var liveGift LiveGift
//	liveGift.SendUid = params.SendId
//	liveGift.SendNickName = sendMember.Nickname
//	liveGift.SendUserImg = sendMember.Avatar
//	liveGift.ReceiveUid = params.ReceiveId
//	liveGift.ReceiveNickName = receiveMember.Nickname
//	liveGift.ReceiveUserImg = receiveMember.Avatar
//	liveGift.GiftId = params.GiftId
//	liveGift.GiftName = giftInfo.Name
//	//蓝钻礼物
//	if giftInfo.Category == 1 {
//		liveGift.GiftPrice = giftInfo.Diamonds
//		liveGift.BallGoldPrice = 0
//		liveGift.GiftType = "蓝钻"
//	} else {
//		liveGift.GiftPrice = 0
//		liveGift.BallGoldPrice = giftInfo.Ballgold
//		liveGift.GiftType = "球豆"
//	}
//	liveGift.GiftCount = params.GiftCount
//	liveGift.RoomId = params.RoomId
//	liveGift.ReceiveCount = giftInfo.Diamonds * params.GiftCount
//	liveGift.GiftIcon = giftInfo.Icon
//	liveGift.GiftEffect = giftInfo.Effect
//	liveGift.GiftEffectImage = giftInfo.EffectImage
//	liveGift.EffectWebp = giftInfo.EffectWebp
//	liveGift.GiftPlayTime = giftInfo.GiftRebateInfo.GiftPlayTime
//	liveGift.GiftTypeId = giftInfo.TypeId
//	liveGift.SendUserLevel = sendMember.Level
//	//发送者等级
//	liveGift.SendUserType = sendMember.Level
//	liveGift.ReceiveUserLevel = receiveMember.Level
//	//接收者等级
//	liveGift.ReceiveUserType = receiveMember.Level
//	//计算礼物收益
//	var exchangeCommission float64
//	//如果是测试账号送的礼物
//	if sendMember.IsStaff == 1 {
//		exchangeCommission = 0
//	} else {
//		if anchorsMember.GuildId > 0 {
//			exchangeCommission = (giftInfo.Diamonds * params.GiftCount) * 0.1 * anchorsMember.GuildDivide * memberUnion.GuildDivide
//		} else {
//			exchangeCommission = (giftInfo.Diamonds * params.GiftCount) * 0.1 * anchorsMember.Divide
//		}
//	}
//	liveGift.ExchangeCommission = exchangeCommission * 100
//	liveGift.GuildDivide = memberUnion.GuildDivide
//	liveGift.Divide = anchorsMember.Divide
//	//如果有工会的情况下，用工会比例
//	if anchorsMember.GuildId > 0 {
//		liveGift.Divide = memberUnion.GuildDivide
//	}
//	liveGift.IsStaff = sendMember.IsStaff
//	liveGift.GiftBackpackId = params.GiftBackpackId
//	liveGift.GiftHotValue = giftInfo.GiftHotValue
//	liveGift.GiftCategory = giftInfo.Category
//	liveGift.GiftSendType = giftInfo.GiftSendType
//	//如果是单聊场景，写入标识
//	if params.SceneType == 1 {
//		liveGift.SceneType = 1
//	}
//
//	jsonBytes, err := json.Marshal(liveGift)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-json格式化失败]-[%s]", err.Error())
//		return
//	}
//
//	//把记录写回mq
//	err = liveCh.ExchangeDeclare(
//		"account-exchang", // name
//		"direct",          // type
//		true,              // durable
//		false,             // auto-deleted
//		false,             // internal
//		false,             // no-wait
//		nil,               // arguments
//	)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-exchange失败]-[%s]", err.Error())
//		return
//	}
//
//	q, err := liveCh.QueueDeclare(
//		initialize.Config.GetString("RabbitMQ.LiveSendQueue"), //name
//		true,  //durable
//		false, //delete when usused
//		false, //exclusive
//		false, //nowait
//		nil,   //argments
//	)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-queue失败]-[%s]", err.Error())
//		return
//	}
//
//	err = liveCh.QueueBind(
//		q.Name,                    // queue name
//		"account-payed-routerkey", // routing key
//		"account-exchang",         // exchange
//		false,
//		nil)
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-绑定queue失败]-[%s]", err.Error())
//		return
//	}
//	err = liveCh.Publish(
//		"account-exchang",         // exchange
//		"account-payed-routerkey", // routing key
//		false,                     // mandatory
//		false,                     // immediate
//		amqp.Publishing{
//			ContentType: "text/json",
//			Body:        jsonBytes,
//		})
//	if err != nil {
//		initialize.IrisLog.Errorf("[推送粉丝经验-发送失败]-[%s]", err.Error())
//		return
//	}
//
//	initialize.IrisLog.Infof("[推送粉丝经验-推送成功]-[%s]", string(jsonBytes))
//}
//
////处理mq商城订单
//func (this *rabbitService) PullShopOrder(number int) {
//	//创建channel
//	shopCh, err := initialize.MqClient.GetChannel()
//	if err != nil {
//		initialize.IrisLog.Infof("[消费商城订单-创建channel失败]-[%s]", err.Error())
//		return
//	}
//	defer shopCh.Close()
//	//队列
//	q, err := shopCh.QueueDeclare(
//		initialize.Config.GetString("RabbitMQ.SendGiftQueue"), //name
//		true,  //durable
//		false, //delete when usused
//		false, //exclusive
//		false, //nowait
//		nil,   //argments
//	)
//	if err != nil {
//		initialize.IrisLog.Errorf("[消费商城订单-创建queue失败]-[%s]", err.Error())
//		return
//	}
//
//	//消费队列
//	msgs, err := shopCh.Consume(
//		q.Name, // queue
//		"",     // consumer
//		false,  // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	if err != nil {
//		initialize.IrisLog.Errorf("[消费商城订单-创建消费者失败]-[%s]", err.Error())
//		return
//	}
//
//	forever := make(chan bool)
//	go func() {
//		for d := range msgs {
//			time.Sleep(20 * time.Millisecond)
//			//开始事务处理各种订单，生成订单，计算收益，生成收益表
//			//绑定参数
//			var query datamodels.SendGiftParams
//			if err := json.Unmarshal(d.Body, &query); err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-json数据失败]-[%s]", err.Error())
//				continue
//			}
//			initialize.IrisLog.Infof("[消费商城订单-mq信息]-[%s]", libs.StructToJson(query))
//
//			//礼物详情
//			giftInfo, _ := this.giftService.Info(query.GiftId)
//
//			//发送礼物用户的信息
//			sendMember, err := this.memberService.SelectInfo(query.SendId)
//			if err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-发送礼物用户异常]-[%s]", err.Error())
//				continue
//			}
//			//发送礼物商城用户信息
//			sendShopMemer, err := this.shopMemberService.SelectInfo(query.SendId)
//			if err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-发送礼物商城用户异常]-[%s]", err.Error())
//				continue
//			}
//			//接收礼物用户的信息
//			receiveMember, err := this.memberService.SelectInfo(query.ReceiveId)
//			if err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-接收礼物用户异常]-[%s]", err.Error())
//				continue
//			}
//			//主播信息
//			anchorsMember, err := this.jmfMemberService.SelectInfo(query.ReceiveId)
//			if err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-主播信息异常异常]-[%s]", err.Error())
//				continue
//			}
//
//			//兑换记录
//			var jmfGiftsLiveExchange datamodels.JmfLiveExchangeRecord
//			jmfGiftsLiveExchange.ID = 0
//			jmfGiftsLiveExchange.OrderId = 0
//			jmfGiftsLiveExchange.GoodsId = giftInfo.Id
//			jmfGiftsLiveExchange.GoodsPrice = giftInfo.Diamonds
//			jmfGiftsLiveExchange.GoodsCount = query.GiftCount
//			jmfGiftsLiveExchange.GoodsIcon = giftInfo.Icon
//			jmfGiftsLiveExchange.RoomId = query.RoomId
//			jmfGiftsLiveExchange.SendId = query.SendId
//			jmfGiftsLiveExchange.SendName = sendMember.Nickname
//			jmfGiftsLiveExchange.ReceiveId = query.ReceiveId
//			jmfGiftsLiveExchange.ReceiveName = receiveMember.Nickname
//
//			//蓝钻礼物
//			if giftInfo.Category == 1 {
//				jmfGiftsLiveExchange.GoodsPrice = giftInfo.Diamonds
//				jmfGiftsLiveExchange.ExchangeType = "diamonds"
//				//如果是背包礼物
//				if query.GiftBackpackId > 0 {
//					jmfGiftsLiveExchange.Total = sendMember.Credit2 + sendMember.Credit5 //当前余额
//				} else {
//					if query.NobleDiamonds > 0 {
//						jmfGiftsLiveExchange.Total = (sendMember.Credit2 + sendMember.Credit5) - (giftInfo.Diamonds * query.GiftCount) //当前用户贵族蓝钻余额
//					}
//					if query.Diamonds > 0 {
//						jmfGiftsLiveExchange.Total = (sendMember.Credit2 + sendMember.Credit5) - (giftInfo.Diamonds * query.GiftCount) //当前用户蓝钻余额
//					}
//					if query.NobleDiamonds > 0 && query.Diamonds > 0 {
//						jmfGiftsLiveExchange.Total = (sendMember.Credit2 + sendMember.Credit5) - (giftInfo.Diamonds * query.GiftCount) //当前用户蓝钻余额
//					}
//				}
//			} else if giftInfo.Category == 2 {
//				jmfGiftsLiveExchange.GoodsPrice = giftInfo.Ballgold
//				jmfGiftsLiveExchange.ExchangeType = "beans"
//				jmfGiftsLiveExchange.Total = sendMember.Credit3 - (giftInfo.Ballgold * query.GiftCount) //当前用户球豆余额
//			} else if giftInfo.Category == 3 {
//				jmfGiftsLiveExchange.GoodsPrice = giftInfo.GiftHotValue //蓝钻或者球豆价格都是热度值
//				jmfGiftsLiveExchange.ExchangeType = "heat"              //热度礼物
//				jmfGiftsLiveExchange.Total = 0                          //热度值
//			}
//
//			//判断是否是奖励的礼物 赠品
//			if giftInfo.Status == 2 {
//				//蓝钻礼物
//				if giftInfo.Diamonds > 0 {
//					jmfGiftsLiveExchange.GoodsPrice = giftInfo.Diamonds
//					jmfGiftsLiveExchange.Total = sendMember.Credit2
//				}
//				//球都礼物
//				if giftInfo.Ballgold > 0 {
//					jmfGiftsLiveExchange.GoodsPrice = giftInfo.Ballgold
//					jmfGiftsLiveExchange.Total = sendMember.Credit3
//				}
//				//jmfGiftsLiveExchange.ExchangeType = "reward" //奖励礼物
//			}
//
//			//如果有工会,先算工会收益，在算个人收益
//			if anchorsMember.GuildId > 0 {
//				memberUnion, err := this.jmfMemberUnionService.SelectInfo(anchorsMember.GuildId)
//				if err != nil {
//					initialize.IrisLog.Errorf("[消费商城订单-查询工会信息失败]-[%s]", err.Error())
//					continue
//				}
//				initialize.IrisLog.Infof("[消费商城订单-工会信息]-[%s]", libs.StructToJson(memberUnion))
//				jmfGiftsLiveExchange.GuildId = memberUnion.GuildId
//				jmfGiftsLiveExchange.GuildName = memberUnion.GuildName
//				jmfGiftsLiveExchange.ExchangeDivide = anchorsMember.GuildDivide
//				jmfGiftsLiveExchange.ExchangeCommission = (giftInfo.Diamonds * query.GiftCount) * 0.1 * memberUnion.GuildDivide * anchorsMember.GuildDivide
//				if anchorsMember.Freeze == 0 {
//					jmfGiftsLiveExchange.GuildCommission = (giftInfo.Diamonds * query.GiftCount) * 0.1 * memberUnion.GuildDivide
//					jmfGiftsLiveExchange.GuildDivide = memberUnion.GuildDivide
//				}else {
//					jmfGiftsLiveExchange.GuildCommission = 0
//					jmfGiftsLiveExchange.GuildDivide = 0
//				}
//			} else {
//				if anchorsMember.Freeze == 0 {
//					jmfGiftsLiveExchange.ExchangeCommission = (giftInfo.Diamonds * query.GiftCount) * 0.1 * anchorsMember.Divide
//					jmfGiftsLiveExchange.ExchangeDivide = anchorsMember.Divide
//				}else {
//					jmfGiftsLiveExchange.ExchangeCommission = 0
//					jmfGiftsLiveExchange.ExchangeDivide = 0
//				}
//			}
//			jmfGiftsLiveExchange.Remark = giftInfo.Name
//			//如果是公司员工送的礼物，工会收益为零，主播收益为零
//			if sendShopMemer.IsStaff == 1 {
//				jmfGiftsLiveExchange.ExchangeCommission = 0
//				jmfGiftsLiveExchange.GuildCommission = 0
//				jmfGiftsLiveExchange.IsStaff = sendShopMemer.IsStaff   //是否公司用户
//				jmfGiftsLiveExchange.IsBorrow = sendShopMemer.IsBorrow //是否借款用户
//			}
//			jmfGiftsLiveExchange.CreateTime = query.SendTime
//			jmfGiftsLiveExchange.RoomName = query.RoomName
//			jmfGiftsLiveExchange.AnchorId = query.AnchorId
//			jmfGiftsLiveExchange.StartTime = query.StartTime
//			jmfGiftsLiveExchange.Platform = query.Platform
//			jmfGiftsLiveExchange.Diamonds = query.Diamonds
//			jmfGiftsLiveExchange.NobleDiamonds = query.NobleDiamonds
//			//主播分类，0是普通主播，1是工会主播，2是员工主播
//			anthorType := 0
//			if anchorsMember.GuildId > 0 {
//				anthorType = 1
//			}
//			if anchorsMember.GuildId > 0 && anchorsMember.GuildName == "官方公会" {
//				anthorType = 2
//			}
//			jmfGiftsLiveExchange.AnchorType = anthorType
//			jmfGiftsLiveExchange.Rate = 0.1
//			jmfGiftsLiveExchange.IsReal = 1
//
//			if query.GiftBackpackId > 0 {
//				jmfGiftsLiveExchange.GiftBackpackId = query.GiftBackpackId
//			}
//
//			/*
//				球豆礼物加入缓存
//			*/
//			var recordInfo datamodels.CreditsRecord
//			recordInfo.ID = 0
//			recordInfo.UID = sendMember.UID
//			recordInfo.Uniacid = 3
//			recordInfo.Credittype = "credit3"
//			recordInfo.Num = -(giftInfo.Ballgold * query.GiftCount)
//			recordInfo.Module = "ewei_shopv2"
//			recordInfo.Createtime = time.Now().Unix()
//			recordInfo.Type = "send_gift"
//			recordInfo.Remark = receiveMember.Nickname + "(" + fmt.Sprintf("%d", jmfGiftsLiveExchange.RoomId) + ")"
//			recordInfo.Total = sendMember.Credit3 - (giftInfo.Ballgold * query.GiftCount)
//
//			jsonBytes, _ := json.Marshal(jmfGiftsLiveExchange)
//			initialize.IrisLog.Infof("[消费商城订单-礼物数据]-[%s]", libs.StructToJson(jmfGiftsLiveExchange))
//			//生产订单和记录
//			logId, err := this.jmfLiveExchangeService.Insert(jmfGiftsLiveExchange, recordInfo)
//			if err != nil {
//				initialize.IrisLog.Infof("[消费商城订单-入库异常]-[%s]", err)
//			}
//			if err != nil {
//				//失败了 记录失败日志
//				var liveLog datamodels.JmfLiveExchangeRecordLog
//				liveLog.ReceiveId = query.ReceiveId
//				liveLog.SendId = query.SendId
//				liveLog.RoomId = query.RoomId
//				liveLog.GoodsId = query.GiftId
//				liveLog.CreateTime = query.SendTime
//				liveLog.Remark = string(jsonBytes)
//				if err := this.jmfLiveExchangeLogService.Insert(liveLog); err != nil {
//					initialize.IrisLog.Errorf("[消费商城订单-插入错误日志失败]-[%s]", string(jsonBytes))
//					continue
//				}
//			}
//
//			//处理返利礼物
//			if giftInfo.GiftRebateInfo.GiftRebateType > 0 {
//				sendGiftKey := repositories.ReturnRedisKey(repositories.API_CACHE_SEND_GIFT_TOTAL, jmfGiftsLiveExchange.SendId) + ":" + time.Now().Format("20060102") + ":" + fmt.Sprintf("%d", jmfGiftsLiveExchange.GoodsId)
//				initialize.RedisCluster.IncrBy(sendGiftKey, int64(jmfGiftsLiveExchange.GoodsCount))
//				initialize.RedisCluster.Expire(sendGiftKey, time.Hour*24)
//				go this.rebateShopGift(jmfGiftsLiveExchange, giftInfo)
//			}
//
//			//日志参数记录更新
//			where := map[string]interface{}{
//				"send_id":    query.SendId,
//				"receive_id": query.ReceiveId,
//				"room_id":    query.RoomId,
//				"send_time":  query.SendTime,
//			}
//			params := map[string]interface{}{
//				"record_id": logId,
//			}
//			if err := this.jmfLiveExchangeParamsService.Update(where, params); err != nil {
//				initialize.IrisLog.Errorf("[消费商城订单-更新礼物日志参数失败]-[%s]", string(jsonBytes))
//			}
//
//			//背包礼物
//			if query.GiftBackpackId > 0 {
//				where := map[string]interface{}{
//					"id":      query.GiftBackpackId,
//					//"status":  0,
//					"user_id": query.SendId,
//				}
//				val := map[string]interface{}{
//					"record_id": logId,
//					//"status":    1,
//				}
//				if err := this.memberBackpackService.Update(where, val); err != nil {
//					initialize.IrisLog.Errorf("[消费商城订单-背包礼物更新失败]-[%s]", val)
//				}
//			}
//
//			initialize.IrisLog.Infof("[消费商城订单-消费成功]-[%d]-[%s]", logId, string(jsonBytes))
//
//			d.Ack(false)
//		}
//	}()
//	initialize.IrisLog.Infof(" [消费商城订单-MQ]-线程[%d]", number)
//	<-forever
//}