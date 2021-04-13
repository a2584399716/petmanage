package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)

func GetGoodsListHandler(c *gin.Context){
	//page, limit := getPageInfo(c)
	page,limit,goodType,nameorBarcode:=getGoodsPageInfo(c)
	goods,count, err := logic.GetGoodsList(page, limit,goodType,nameorBarcode)

	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,goods)

}

func GetGoodsTypeListHandler(c *gin.Context){
	page, limit := getPageInfo(c)
	types,count,err := logic.GetGoodsTypeList(page, limit)
	if err != nil {
		zap.L().Error("logic.GetGoodsTypeList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,types)


}

func GetSupplierHandler(c *gin.Context){
	page, limit := getPageInfo(c)
	suppliers,count,err :=logic.GetSupplierList(page, limit)
	if err != nil {
		zap.L().Error("logic.GetSupplierList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,suppliers)
}

func GetSuitableTypeHandler(c *gin.Context){
	page, limit := getPageInfo(c)
	Suitables,count,err :=logic.GetSuitableType(page, limit)
	if err != nil {
		zap.L().Error("logic.GetSuitableType() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,Suitables)
}


//AddGoodsHandler 添加商品
func AddGoodsHandler(c *gin.Context){
	g := new(models.Good)
	if err := c.ShouldBind(g); err != nil {
		zap.L().Debug("c.ShouldBindJSON(g) error", zap.Any("err", err))
		zap.L().Error("add goods with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增商品
	if err := logic.AddGoods(g); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddGoods(g) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}

func AddSupplierHandler(c *gin.Context){
	sup := new(models.Supplier)
	if err := c.ShouldBind(sup); err != nil {
		zap.L().Debug("c.ShouldBindJSON(sup) error", zap.Any("err", err))
		zap.L().Error("add Supplier with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增供应商
	if err := logic.AddSupplier(sup); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddGoods(g) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}

//AddGoodsTypeHandler 添加商品类型
func AddGoodsTypeHandler(c *gin.Context){
	t := new(models.GoodType)
	if err := c.ShouldBind(t); err != nil {
		zap.L().Debug("c.ShouldBindJSON(g) error", zap.Any("err", err))
		zap.L().Error("add goodstype with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增商品
	if err := logic.AddGoodsType(t); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddGoodsType(t) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}
//AddSuitableTypeHandler 新增商品适用类型
func AddSuitableTypeHandler(c *gin.Context){
	s := new(models.SuitableType)
	if err := c.ShouldBind(s); err != nil {
		zap.L().Debug("c.ShouldBindJSON(s) error", zap.Any("err", err))
		zap.L().Error("add SuitableType with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增适用类型
	if err := logic.AddSuitableType(s); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddSuitableType(s) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}

//EditGoodsHandlerv 修改商品
func EditGoodsHandler(c *gin.Context){
	g := new(models.Good)
	if err := c.ShouldBind(g); err != nil {
		zap.L().Debug("c.ShouldBindJSON(g) error", zap.Any("err", err))
		zap.L().Error("edit goods with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//修改商品
	if err := logic.EditGoods(g); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.EditGoods(g) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)


}

//DelGoodsByIdHandler 删除指定商品
func DelGoodsByIdHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.DelGoodsById(id); err != nil {

		zap.L().Error("logic.DelGoodsTypeById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)

}

//DelGoodsTypeByIdHandler 删除指定商品类型
func DelGoodsTypeByIdHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := int8(id64)
	if err := logic.DelGoodsTypeById(id); err != nil {

		zap.L().Error("logic.DelGoodsTypeById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)

}

func DelSuitableTypeByIdHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := int8(id64)
	if err := logic.DelSuitableTypeById(id); err != nil {

		zap.L().Error("logic.DelSuitableTypeById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}

//DelSupplierByIdHandler 删除指定供应商
func DelSupplierByIdHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := int8(id64)
	if err := logic.DelSupplierById(id); err != nil {

		zap.L().Error("logic.DelSupplierById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}

func EditSupplierHandler(c *gin.Context){
	sup := new(models.Supplier)
	if err := c.ShouldBind(sup); err != nil {
		zap.L().Debug("c.ShouldBindJSON(sup) error", zap.Any("err", err))
		zap.L().Error("edit Supplier with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//修改供应商
	if err := logic.EditSupplier(sup); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.EditSupplier(sup) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}