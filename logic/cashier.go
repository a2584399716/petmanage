package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCashierList() (cashiers []*models.CashierDetail, err error) {
	//得到商品详情

	cashier1, err := GetCashierGoodsList()
	//加入到商品结构体
	cashier2, err := GetCashierServiceList()
	cashier3, err := GetCashierLivingList()
	cashiers = append(cashiers, cashier1)
	cashiers = append(cashiers, cashier2)
	cashiers = append(cashiers, cashier3)
	return
}

func GetCashierGoodsList() (cashier *models.CashierDetail, err error) {
	cashier = new(models.CashierDetail)
	//得到商品类型名
	cashier.Cashier.Name = "商品"
	//得到商品类型详情
	var types []*models.GoodType
	types, _, err = GetGoodsTypeList(1, 20)
	for _, typeDetail := range types {
		chiledd := new(models.Children)
		good_type_name, err := mysql.GetGoodTypeNameByTypeID(typeDetail.TypeId)
		if err != nil {
			zap.L().Error("mysql.GetGoodTypeNameByTypeID(typeDetail.TypeId) failed",
				zap.Int8("TypeId", typeDetail.TypeId),
				zap.Error(err))
			continue
		}

		chiledd.Name = good_type_name
		chiledd.Id = int64(typeDetail.TypeId)

		goods, _, err := mysql.GetGoodListByGoodType(1, 20, typeDetail.TypeId)
		for _, good := range goods {
			childDetail := &models.Cashier{
				Id:    good.Good_id,
				Name:  good.Good_name,
				Price: good.Selling_Price,
				End:   true,
			}
			chiledd.Child = append(chiledd.Child, childDetail)
		}
		cashier.Chiledd = append(cashier.Chiledd, chiledd)
	}
	//返回详情

	return
}

func GetCashierServiceList() (cashier *models.CashierDetail, err error) {
	cashier = new(models.CashierDetail)
	//得到商品类型名
	cashier.Cashier.Name = "服务"
	//得到商品类型详情
	var types []*models.ServeType
	types, _, err = GetServeTypeList(1, 20)
	for _, typeDetail := range types {
		chiledd := new(models.Children)
		serve_type_name, err := mysql.GetServeTypeNameByTypeID(typeDetail.TypeId)
		if err != nil {
			zap.L().Error("mysql.GetServeTypeNameByTypeID(typeDetail.TypeId) failed",
				zap.Int8("TypeId", typeDetail.TypeId),
				zap.Error(err))
			continue
		}

		chiledd.Name = serve_type_name
		chiledd.Id = int64(typeDetail.TypeId)

		services, err := mysql.GetServiceListByServiceType(typeDetail.TypeId)
		for _, service := range services {
			childDetail := &models.Cashier{
				Id:    int64(service.ServeId),
				Name:  service.ServeName,
				Price: service.ServePrice,
				End:   true,
			}
			chiledd.Child = append(chiledd.Child, childDetail)
		}
		cashier.Chiledd = append(cashier.Chiledd, chiledd)
	}
	//返回详情

	return
}

func GetCashierLivingList() (cashier *models.CashierDetail, err error) {
	cashier = new(models.CashierDetail)
	//得到商品类型名
	cashier.Cashier.Name = "活体"
	//得到商品类型详情
	var types []*models.Varietie
	types, _, err = GetVarietieList(1, 20)
	for _, typeDetail := range types {
		chiledd := new(models.Children)
		living_type_name, err := mysql.GetVarietieNameByVarietieID(typeDetail.VarietieId)
		if err != nil {
			zap.L().Error("mysql.GetVarietieNameByVarietieID(typeDetail.VarietieId) failed",
				zap.Int8("TypeId", typeDetail.VarietieId),
				zap.Error(err))
			continue
		}

		chiledd.Name = living_type_name
		chiledd.Id = int64(typeDetail.VarietieId)

		livings, err := mysql.GetLivingListByVarietieType(typeDetail.VarietieId)
		for _, living := range livings {
			childDetail := &models.Cashier{
				Id:    int64(living.LivingId),
				Name:  living.LivingName,
				Price: living.LivingPrice,
				End:   true,
			}
			chiledd.Child = append(chiledd.Child, childDetail)
		}
		cashier.Chiledd = append(cashier.Chiledd, chiledd)
	}
	//返回详情

	return
}

func GetCustomerInfoByInfo(info string) (data []*models.Customer, err error) {
	data, err = mysql.GetCustomerInfoByInfo(info)

	return
}


func GetCustomerInfoById(id int64)(info *models.UserInfo,err error){
	info,err = mysql.GetCustomerInfoById(id)
	if err != nil {
		zap.L().Error("mysql.GetCustomerInfoById(id) failed",
			zap.Int64("info", id),
			zap.Error(err))
	}
	return
}