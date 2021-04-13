package models

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ParamGoodList struct{
	GoodID int64  `json:"good_id" form:"good_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Limit        int64  `json:"limit" form:"size" example:"10"`      // 每页数据量
	//Order       string `json:"order" form:"order" example:"score"` // 排序依据

}
