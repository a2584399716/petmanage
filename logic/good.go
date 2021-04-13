package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetGoodsList(page int64, limit int64, goodType int8, nameorBarcode string) (data []*models.ApiGoodDetail, count int, err error) {
	goods := make([]*models.Good, 0, 2)
	//1只有page和limit
	if (goodType == 0) && (nameorBarcode == "") {
		goods, count, err = mysql.GetGoodList(page, limit)
	} else if (goodType != 0) && (nameorBarcode != "") {
		//两个都有
		goods, count, err = mysql.GetGoodListByGoodTypeAndNameOrBarcode(page, limit, goodType, nameorBarcode)
	} else if nameorBarcode != "" {
		//只有nameorBarcode ,good_type为0
		goods, count, err = mysql.GetGoodListByNameOrBarcode(page, limit, nameorBarcode)
	} else {
		//只有good_type,nameorBarcode为空
		goods, count, err = mysql.GetGoodListByGoodType(page, limit, goodType)
	}
	if err != nil {
		return nil, 0, err
	}
	data = make([]*models.ApiGoodDetail, 0, len(goods))
	for _, good := range goods {
		good_type_name, err := mysql.GetGoodTypeNameByTypeID(good.Good_Type)
		if err != nil {
			zap.L().Error("mysql.GetGoodTypeNameByTypeID(good.Good_Type) failed",
				zap.Int8("Good_type_id", good.Good_Type),
				zap.Error(err))
			continue
		}
		goodDetail := &models.ApiGoodDetail{
			Good_type_name: good_type_name,
			Good:           good,
		}
		data = append(data, goodDetail)
	}

	return
}

func GetGoodsTypeList(page, limit int64) (data []*models.GoodType, count int, err error) {
	data, count, err = mysql.GetGoodsTypeList(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return
}

func GetSupplierList(page, limit int64) (data []*models.Supplier, count int, err error) {
	data, count, err = mysql.GetSupplierList(page, limit)
	//fmt.Println((&data[0].CreateTime).Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, 0, err
	}
	return
}

func GetSuitableType(page, limit int64) (data []*models.SuitableType, count int, err error) {
	data, count, err = mysql.GetSuitableType(page, limit)
	//fmt.Println((&data[0].CreateTime).Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, 0, err
	}
	return
}

func AddGoods(g *models.Good) (err error) {
	//查询商品名是否重复
	if err = mysql.SearchPepeatByGoodsName(g.Good_name, g.Good_id); err != nil {
		return
	}
	err = mysql.AddGoods(g)
	if err != nil {
		return err
	}

	return
}

func AddSupplier(sup *models.Supplier) (err error) {
	//查询供应商名是否重复
	if err = mysql.SearchPepeatBySupplierName(sup.SupplierName, sup.SupplierId); err != nil {
		return
	}
	err = mysql.AddSupplier(sup)
	if err != nil {
		return err
	}

	return
}

func AddGoodsType(t *models.GoodType) (err error) {
	//查询类型名是否重复
	if err = mysql.SearchPepeatByGoodsTypeName(t.TypeName); err != nil {
		return
	}
	err = mysql.AddGoodsType(t)
	if err != nil {
		return err
	}

	return
}

func AddSuitableType(s *models.SuitableType) (err error) {
	//查询适用类型名是否重复
	if err = mysql.SearchPepeatBySuitableTypeName(s.SuitableName); err != nil {
		return
	}
	err = mysql.AddSuitableType(s)
	if err != nil {
		return err
	}

	return
}

func EditGoods(g *models.Good) (err error) {
	if err = mysql.SearchPepeatByGoodsName(g.Good_name, g.Good_id); err != nil {
		return
	}
	err = mysql.EditGoods(g)
	if err != nil {
		return err
	}

	return
}

func EditSupplier(sup *models.Supplier) (err error) {
	if err = mysql.SearchPepeatBySupplier(sup.SupplierName, sup.SupplierId); err != nil {
		return
	}
	err = mysql.EditSupplier(sup)
	if err != nil {
		return err
	}

	return
}

func DelGoodsById(id int64) (err error) {
	err = mysql.DelGoodsById(id)
	if err != nil {
		zap.L().Error("DelGoodsById(id int64) failed",
			zap.Int64("id", id),
			zap.Error(err))
	}
	return
}

func DelGoodsTypeById(id int8) (err error) {
	err = mysql.DelGoodsTypeById(id)
	if err != nil {
		zap.L().Error("DelGoodsTypeById(id int8) failed",
			zap.Int8("id", id),
			zap.Error(err))
	}
	return
}

func DelSuitableTypeById(id int8) (err error) {
	err = mysql.DelSuitableTypeById(id)
	if err != nil {
		zap.L().Error("DelSuitableTypeById(id int8) failed",
			zap.Int8("id", id),
			zap.Error(err))
	}
	return
}

func DelSupplierById(id int8) (err error) {
	err = mysql.DelSupplierById(id)
	if err != nil {
		zap.L().Error("DelSupplierById(id int8) failed",
			zap.Int8("id", id),
			zap.Error(err))
	}
	return
}
