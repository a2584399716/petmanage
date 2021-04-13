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

func GetVarietieListHandler(c *gin.Context) {
	page, limit := getPageInfo(c)
	varieties, count, err := logic.GetVarietieList(page, limit)
	if err != nil {
		zap.L().Error("logic.GetVarietieList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c, count, varieties)
}

func GetLivingListHandler(c *gin.Context) {
	page, limit := getPageInfo(c)
	livings, count, err := logic.GetLivingList(page, limit)
	if err != nil {
		zap.L().Error("logic.GetLivingList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c, count, livings)
}

//AddVarietieHandler 新增品种
func AddVarietieHandler(c *gin.Context) {
	v := new(models.Varietie)

	if err := c.ShouldBind(v); err != nil {
		zap.L().Debug("c.ShouldBindJSON(sup) error", zap.Any("err", err))
		zap.L().Error("add Varietie with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增供应商
	if err := logic.AddVarietie(v); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddVarietie(v) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func DelVarietieHandler(c *gin.Context) {
	idStr := c.PostForm("id")
	id64, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id := int8(id64)
	if err := logic.DelVarietieById(id); err != nil {

		zap.L().Error("logic.DelVarietieById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func DelLivingByIdHandler(c *gin.Context){
	idStr := c.PostForm("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := logic.DelLivingById(id); err != nil {

		zap.L().Error("logic.DelLivingById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func AddLivingHandler(c *gin.Context) {
	l := new(models.LivingAnimal)
	if err := c.ShouldBind(l); err != nil {
		zap.L().Debug("c.ShouldBindJSON(l) error", zap.Any("err", err))
		zap.L().Error("add Living with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//新增供应商
	if err := logic.AddLiving(l); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.AddLiving(l) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

//EditLivingHandler 修改活体属性
func EditLivingByIdHandler(c *gin.Context) {
	g := new(models.LivingAnimal)
	if err := c.ShouldBind(g); err != nil {
		zap.L().Debug("c.ShouldBindJSON(g) error", zap.Any("err", err))
		zap.L().Error("edit Living with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//修改商品
	if err := logic.EditLivingById(g); err != nil {
		if errors.Is(err, mysql.ErrorGoodsNameExist) {
			ResponseError(c, CodeGoodsNameExist)
			return
		}
		zap.L().Error("logic.EditLivingById(g) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)

}
