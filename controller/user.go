package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

func GetUserListHandler(c *gin.Context){
	page, limit := getPageInfo(c)
	users,count, err := logic.GetUserList(page, limit)

	if err != nil {
		zap.L().Error("logic.GetUserList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithCount(c,count,users)
}
