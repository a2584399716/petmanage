package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetServeList(page, limit int64) (data []*models.ApiServeDetail, count int, err error) {
	serves, count, err := mysql.GetServeList(page, limit)
	if err != nil {
		return nil, 0, err
	}
	data = make([]*models.ApiServeDetail, 0, len(serves))
	for _, serve := range serves {
		serve_type_name, err := mysql.GetServeTypeNameByTypeID(serve.ServeType)
		if err != nil {
			zap.L().Error("mysql.GetServeTypeNameByTypeID(serve.ServeType) failed",
				zap.Int8("serve_type", serve.ServeType),
				zap.Error(err))
			continue
		}
		serveDetail := &models.ApiServeDetail{
			Serve_type_name:serve_type_name,
			Serve:serve,
		}
		data = append(data,serveDetail)
	}
	return
}

func GetServeTypeList(page, limit int64) (data []*models.ServeType, count int, err error) {
	data, count, err = mysql.GetServeTypeList(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return
}


func AddService(s *models.Serve)(err error){
	//查询供应商名是否重复
	if err = mysql.SearchPepeatByServiceName(s.ServeName,s.ServeId); err != nil {
		return
	}
	err = mysql.AddService(s)
	if err != nil {
		return err
	}

	return
}

func EditServiceById(g *models.Serve)(err error){
	if err = mysql.SearchPepeatByServiceName(g.ServeName,g.ServeId); err != nil {
		return
	}
	err = mysql.EditServiceById(g)
	if err != nil {
		return err
	}

	return
}

func AddServiceType(s *models.ServeType)(err error){
	//查询供应商名是否重复
	if err = mysql.SearchPepeatByServiceTypeName(s.TypeName,s.TypeId); err != nil {
		return
	}
	err = mysql.AddServiceType(s)
	if err != nil {
		return err
	}

	return
}



func DelServiceById(id int64)(err error){
	err = mysql.DelServiceById(id)
	if err != nil {
		zap.L().Error("DelServiceById(id int64) failed",
			zap.Int64("id", id),
			zap.Error(err))
	}
	return
}

func DelServiceTypeById(id int8)(err error){
	err = mysql.DelServiceTypeById(id)
	if err != nil {
		zap.L().Error("DelServiceTypeById(id int8) failed",
			zap.Int8("id", id),
			zap.Error(err))
	}
	return
}