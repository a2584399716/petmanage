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

func GetServeListHander(c *gin.Context){
	page, limit := getPageInfo(c)
	serves,count,err :=logic.GetServeList(page, limit)

	if err != nil {
		zap.L().Error("logic.GetServeList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,serves)
}

func AddServiceHandler(c *gin.Context){
	s := new(models.Serve)
	if err := c.ShouldBind(s); err != nil {
		zap.L().Debug("c.ShouldBindJSON(sup) error", zap.Any("err", err))
		zap.L().Error("add service with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}


	//新增供应商
	if err := logic.AddService(s); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddService(s) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}

func EditServiceByIdHandler(c *gin.Context){
	g := new(models.Serve)
	if err := c.ShouldBind(g); err != nil {
		zap.L().Debug("c.ShouldBindJSON(g) error", zap.Any("err", err))
		zap.L().Error("edit serve with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//修改商品
	if err := logic.EditServiceById(g); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.EditServiceById(g) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)


}

func AddServiceTypeHandler(c *gin.Context){
	s := new(models.ServeType)

	if err := c.ShouldBind(s); err != nil {
		zap.L().Debug("c.ShouldBindJSON(sup) error", zap.Any("err", err))
		zap.L().Error("add ServiceType with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//新增供应商
	if err := logic.AddServiceType(s); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddServiceType(s) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}



	ResponseSuccess(c,nil)
}



func GetServeTypeListHandler(c *gin.Context){
	page, limit := getPageInfo(c)
	serveTypes,count,err :=logic.GetServeTypeList(page, limit)
	if err != nil {
		zap.L().Error("logic.GetServeTypeList(page, limit) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,serveTypes)
}

func DelServiceHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.DelServiceById(id); err != nil {

		zap.L().Error("logic.DelServiceById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}

func DelServiceTypeHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := int8(id64)
	if err := logic.DelServiceTypeById(id); err != nil {

		zap.L().Error("logic.DelServiceTypeById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c,nil)
}
