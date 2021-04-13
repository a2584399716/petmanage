package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/controller"
	"web_app/logger"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Static("/static", "./static")

	r.POST("/login", controller.LoginHandler)
	r.LoadHTMLFiles("./static/login.html")
	//r.Use(controller.Auth())
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,"hello")
	//})

	//获取商品列表
	r.GET("/getgoodslist", controller.GetGoodsListHandler)

	//删除指定商品
	r.POST("/delgoodsbyid", controller.DelGoodsByIdHandler)

	//添加商品
	r.POST("/addgoods", controller.AddGoodsHandler)

	//修改商品
	r.POST("/editgoods", controller.EditGoodsHandler)

	//获取商品类型
	r.GET("/getgoodstype", controller.GetGoodsTypeListHandler)

	//添加商品类型
	r.POST("/addgoodstype", controller.AddGoodsTypeHandler)

	//删除指定商品类型
	r.POST("/delgoodstypebyid", controller.DelGoodsTypeByIdHandler)

	//获取供应商列表
	r.GET("/getsupplier", controller.GetSupplierHandler)

	//新增供应商
	r.POST("/addsupplier",controller.AddSupplierHandler)

	//修改供应商
	r.POST("/editsupplier",controller.EditSupplierHandler)

	//删除指定供应商
	r.POST("/delsupplierbyid",controller.DelSupplierByIdHandler)

	//获取商品适用类型
	r.GET("/getsuitabletype", controller.GetSuitableTypeHandler)

	//新增商品适用类型
	r.POST("addsuitablestype",controller.AddSuitableTypeHandler)

	//删除商品适用类型
	r.POST("/delsuitabletypebyid",controller.DelSuitableTypeByIdHandler)

	//获取服务列表
	r.GET("/getservelist", controller.GetServeListHander)

	//新增服务
	r.POST("/addservice",controller.AddServiceHandler)

	//删除指定服务
	r.POST("/delservice",controller.DelServiceHandler)

	//修改指定服务
	r.POST("/editservice",controller.EditServiceByIdHandler)

	//获取服务类型列表
	r.GET("/getservetypelist", controller.GetServeTypeListHandler)

	//新增服务类型
	r.POST("/addservicetype",controller.AddServiceTypeHandler)

	//删除指定服务类型
	r.POST("/delserviceType",controller.DelServiceTypeHandler)

	//获取宠物种类列表
	r.GET("/getvarietielist", controller.GetVarietieListHandler)

	//新增种类品种
	r.POST("/addvarietie",controller.AddVarietieHandler)

	//删除宠物种类
	r.POST("/delvarietie",controller.DelVarietieHandler)

	//获取活体列表
	r.GET("/getlivinglist", controller.GetLivingListHandler)

	//新增活体
	r.POST("/addliving",controller.AddLivingHandler)

	//修改活体属性
	r.POST("/editliving",controller.EditLivingByIdHandler)


	//删除活体
	r.POST("/delliving",controller.DelLivingByIdHandler)

	//获取收银台信息
	r.POST("/getcashierlist",controller.GetCashierListHandler)

	//收银时获取会员信息
	r.GET("/getcustomerinfo",controller.GetCustomerInfoHandler)

	//收银时通过会员id获取相关信息
	r.GET("/getcustomerinfobyid",controller.GetCustomerInfoByIdHandler)


	//获取会员列表
	r.GET("/getuserlist", controller.GetUserListHandler)

	return r
}
