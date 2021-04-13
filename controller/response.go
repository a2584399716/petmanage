package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code":10001,   //程序中的错误码
	"msg":xx,		//提示信息
	"data":{}		//数据
}

*/

type ResponseData struct {
	Code ResCode         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context,code ResCode){
	rd :=&ResponseData{
		Code: code,
		Msg:  getMsg(code),
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}

func ResponseErrorWithMsg(c *gin.Context,code ResCode,msg interface{}){
	rd :=&ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}

func ResponseSuccess(c *gin.Context,data interface{}){
	rd :=&ResponseData{
		Code: CodeSuccess,
		Msg:  getMsg(CodeSuccess),
		Data: data,
	}
	c.JSON(http.StatusOK,rd)
}
func ResponseSuccessWithCount(c *gin.Context,count int,data interface{}){

	c.JSON(http.StatusOK,	gin.H{
		"code": 0,
		"count":count,
		"msg":  "",
		"data": data,
	})
}

