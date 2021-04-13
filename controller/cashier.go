package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"web_app/logic"
)

func GetCashierListHandler(c *gin.Context) {
	//cashiers := make([]*models.CashierDetail, 0, 2)

	cashiers, err := logic.GetCashierList()
	if err != nil {
		zap.L().Error("logic.GetCashierList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	c.JSON(http.StatusOK, cashiers)

}

func GetCustomerInfoHandler(c *gin.Context) {
	info := c.Query("info")
	data, err := logic.GetCustomerInfoByInfo(info)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetCustomerInfoByIdHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {

		zap.L().Error("get userid with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	info,err :=logic.GetCustomerInfoById(id)
	if err != nil {
		zap.L().Error("logic.GetCustomerInfoById(id) failed",
			zap.Int64("id", id),
			zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c,info)
}
