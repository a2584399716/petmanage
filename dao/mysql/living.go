package mysql

import (
	"web_app/models"
)

func GetVarietieList(page, limit int64) (varietie []*models.Varietie, count int, err error) {
	sqlStr := `select varietie_id,varietie_name,varietie_number from varietie ORDER BY varietie_id ASC limit ?,?`
	sqlStr2 := `select count(*) from varietie`
	varietie = make([]*models.Varietie, 0, 2)
	err = db.Select(&varietie, sqlStr, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2)
	return
}

func SearchPepeatByvarietieName(name string, id int8) (err error) {
	sqlStr := `select count(*) from varietie where varietie_name = ? and varietie_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func SearchPepeatByLivingName(name string, id int64) (err error) {
	sqlStr := `select count(*) from living_animal where living_name = ? and living_id <> ? limit 1`
	var count int
	err = db.Get(&count, sqlStr, name, id)
	if count != 0 {
		err = ErrorGoodsNameExist
	}
	return
}

func AddVarietie(s *models.Varietie) (err error) {
	sqlStr := `insert into varietie(varietie_name)
				values (?)
	`
	_, err = db.Exec(sqlStr, s.VarietieName)
	return
}

func AddLiving(l *models.LivingAnimal) (err error) {
	sqlStr := `insert into living_animal(
				living_name, varietie_id, living_sex,cost_price,selling_price,living_state)
				values (?, ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, l.LivingName, l.VarietieId, l.LivingSex, l.CostPrice, l.LivingPrice, l.LivingState)
	return
}

func EditLivingById(g *models.LivingAnimal) (err error) {
	sqlStr := `update living_animal set living_name=?,varietie_id=?,living_sex=?,cost_price=?,selling_price=?,living_state=? where living_id = ?`
	_, err = db.Exec(sqlStr, g.LivingName, g.VarietieId, g.LivingSex, g.CostPrice, g.LivingPrice, g.LivingState, g.LivingId)
	return
}

func GetLivingList(page, limit int64) (livings []*models.LivingAnimal, count int, err error) {
	sqlStr := `select 
	living_id, living_name, varietie_id, living_sex, cost_price,selling_price,living_state,create_time
	from living_animal
	ORDER BY create_time
	DESC
	limit ?,?
	`
	sqlStr2 := `select count(*) from living_animal`
	livings = make([]*models.LivingAnimal, 0, 2)
	err = db.Select(&livings, sqlStr, (page-1)*limit, limit)
	err = db.Get(&count, sqlStr2)
	return
}

func GetVarietieNameByVarietieID(id int8) (varietieName string, err error) {
	sqlStr := `select varietie_name from varietie where varietie_id = ?`
	err = db.Get(&varietieName, sqlStr, id)
	return
}

func GetLivingSexNameBySexID(id int8) (sexName string, err error) {
	sqlStr := `select sex_name from sex where sex_id = ?`
	err = db.Get(&sexName, sqlStr, id)
	return
}

func GetStateNameByStateID(id int8) (sexName string, err error) {
	sqlStr := `select state_name from state where state_id = ?`
	err = db.Get(&sexName, sqlStr, id)
	return
}

func DelVarietieById(id int8) (err error) {
	sqlStr := `delete from varietie where varietie_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}

func DelLivingById(id int64) (err error) {
	sqlStr := `delete from living_animal where living_id = ?`
	_, err = db.Exec(sqlStr, id)
	return
}
