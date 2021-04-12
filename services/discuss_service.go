package services

import (
	"bbs/repositories"
)

/*
	提供关于讨论帖子服务

	作者名称：leixiaotian 创建时间：20210412
*/
type DiscussInterfaceService interface {
	GetWallet(userId int64) (interface{},error) //获取用户钱包
}

//初始化对象函数
func NewDiscussService() DiscussInterfaceService {
	return &discussService{
		shopMemberService:       repositories.NewBbsDiscuss(),
	}
}

type discussService struct {
	shopMemberService 			repositories.BbsDiscussInterface     //商城会员服务
}

//获取用户钱包
func (this *discussService) GetWallet(userId int64) (interface{},error){

	type UserWalletRow struct {
		ID                int     `json:"id"` 					 //ID
		UserId            int64   `json:"user_id"`               //用户中心id
		TotalCommission   float64 `json:"total_commission"`      //历史佣金，总佣金
		CurrentCommission float64 `json:"current_commission"`    //当前佣金
		DrawCommission    float64 `json:"draw_commission"`       //已提取佣金
		CreateTime        int64   `json:"create_time"`           //创建时间
		UpdateTime        int64   `json:"update_time"`           //创建时间
		ShowCommission    float64 `json:"show_commission"`       //前端显示佣金
	}
	var expertWalletRow UserWalletRow
	//expertWallet, err := this.jmfUserWalletService.GetUserWalletInfo(userId)
	//if err != nil {
	//	return 3006, err
	//}
	expertWalletRow.ID = 1
	expertWalletRow.UserId = 222
	//expertWalletRow.TotalCommission = expertWallet.TotalCommission
	//expertWalletRow.CurrentCommission = expertWallet.CurrentCommission
	//expertWalletRow.DrawCommission = expertWallet.DrawCommission
	//expertWalletRow.CreateTime = expertWallet.CreateTime
	//expertWalletRow.UpdateTime = expertWallet.UpdateTime
	//currentCommissionFloor := math.Floor(expertWallet.CurrentCommission * 100)
	//expertWalletRow.ShowCommission, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", currentCommissionFloor * 0.01), 64)
	return expertWalletRow, nil
}
