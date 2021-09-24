package controllers

import (
	"bbs/initialize"
	"bbs/libs"
	"bbs/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GoodsController struct {

	GoodsService services.GoodsInterfaceService
	RpcGoodsService services.GoodsServiceClient
}

func NewGoodsController() *GoodsController {
	obj := &GoodsController{
		GoodsService: services.NewGoodsService(),
		RpcGoodsService: services.NewGoodsServiceClient(initialize.GrpcConn),
	}
	return obj
}


//获取用户详情页
func (this *GoodsController) Detail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.IrisLog.Infof("[用户控制器-http请求数据]-[%d]", id)

	//goodsClient := services.NewGoodsServiceClient(initialize.GrpcConn) //优化如下
	//goodsRes, err := goodsClient.GetGoodsStock(context.Background(), &services.GoodsRequest{GoodsId:2}) //优化如下

	goodsRes, err := this.RpcGoodsService.GetGoodsStock(context.Background(), &services.GoodsRequest{GoodsId: int32(id)})
	if err != nil {
		initialize.IrisLog.Errorf("[grpc请求报错：]-[%s]", err)
		c.JSON(http.StatusInternalServerError,  libs.ReturnJson(500, "", gin.H{}))
		return
	}

	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", result))
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"goods_id": id, "remark": "库存数量等于goods_id乘以10", "goods": goodsRes}))

	//defer  initialize.GrpcConn.Close() //不能关闭，关闭第二次请求报链接关闭错误
}