package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	var (
		page  int64
		limit int64
		err   error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return page, limit
}

func getGoodsPageInfo(c *gin.Context) (int64, int64, int8, string) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	goodTypeStr := c.Query("good_type")
	nameOrBarCodeStr := c.Query("name_barcode")

	var (
		page          int64
		limit         int64
		goodType      int8
		err           error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	goodType64, err := strconv.ParseInt(goodTypeStr, 10, 8)
	goodType = int8(goodType64)
	if err != nil {
		goodType = 0
	}



	return page, limit,goodType,nameOrBarCodeStr
}
