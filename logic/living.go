package logic

import (
	"fmt"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetVarietieList(page, limit int64) (data []*models.Varietie, count int, err error) {
	data, count, err = mysql.GetVarietieList(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return
}

func GetLivingList(page, limit int64) (data []*models.ApiLivingDetail, count int, err error) {
	livings, count, err := mysql.GetLivingList(page, limit)
	fmt.Println(livings)
	if err != nil {
		return nil, 0, err
	}
	data = make([]*models.ApiLivingDetail, 0, len(livings))
	for _, living := range livings {
		varietieName, err := mysql.GetVarietieNameByVarietieID(living.VarietieId)
		if err != nil {
			zap.L().Error("GetVarietieNameByVarietieID(living.VarietieId) failed",
				zap.Int8("VarietieId", living.VarietieId),
				zap.Error(err))
			continue
		}

		sexName, err := mysql.GetLivingSexNameBySexID(living.LivingSex)
		if err != nil {
			zap.L().Error("mysql.GetLivingSexNameBySexID(living.LivingSex) failed",
				zap.Int8("LivingSexId", living.LivingSex),
				zap.Error(err))
			continue
		}

		stateName, err := mysql.GetStateNameByStateID(living.LivingState)
		if err != nil {
			zap.L().Error("mysql.GetStateNameByStateID(living.LivingState) failed",
				zap.Int8("LivingStateID", living.LivingState),
				zap.Error(err))
			continue
		}

		livingDetail := &models.ApiLivingDetail{
			VarietieName: varietieName,
			SexName:      sexName,
			StateName:    stateName,
			LivingAnimal: living,
		}
		data = append(data, livingDetail)
	}
	return
}

func AddVarietie(s *models.Varietie) (err error) {
	//查询供应商名是否重复
	if err = mysql.SearchPepeatByvarietieName(s.VarietieName, s.VarietieId); err != nil {
		return
	}
	err = mysql.AddVarietie(s)
	if err != nil {
		return err
	}

	return
}

func DelVarietieById(id int8) (err error) {
	err = mysql.DelVarietieById(id)
	if err != nil {
		zap.L().Error("DelVarietieById(id int8) failed",
			zap.Int8("id", id),
			zap.Error(err))
	}
	return
}

func DelLivingById(id int64) (err error) {
	err = mysql.DelLivingById(id)
	if err != nil {
		zap.L().Error("DelLivingById(id int64) failed",
			zap.Int64("id", id),
			zap.Error(err))
	}
	return
}

func AddLiving(l *models.LivingAnimal) (err error) {
	//查询供应商名是否重复
	if err = mysql.SearchPepeatByLivingName(l.LivingName, l.LivingId); err != nil {
		return
	}
	err = mysql.AddLiving(l)
	if err != nil {
		return err
	}

	return
}

func EditLivingById(g *models.LivingAnimal) (err error) {
	if err = mysql.SearchPepeatByLivingName(g.LivingName, g.LivingId); err != nil {
		return
	}
	err = mysql.EditLivingById(g)
	if err != nil {
		return err
	}

	return
}
