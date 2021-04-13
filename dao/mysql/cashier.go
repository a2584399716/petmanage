package mysql

import "web_app/models"

func GetGoodsCashier(cashierDetail *models.CashierDetail) {

}

func GetServiceListByServiceType(serviceType int8) (services []*models.Serve, err error) {
	sqlStr := `SELECT serve_id, serve_name, selling_price FROM serve 
				where serve_type = ?
				ORDER BY create_time
				DESC
`

	services = make([]*models.Serve, 0, 2)
	err = db.Select(&services, sqlStr, serviceType)

	return
}

func GetLivingListByVarietieType(varietieId int8) (livings []*models.LivingAnimal, err error) {
	sqlStr := `SELECT living_id, living_name, selling_price FROM living_animal 
				where varietie_id = ? and living_state = 1
				ORDER BY create_time
				DESC
`

	livings = make([]*models.LivingAnimal, 0, 2)
	err = db.Select(&livings, sqlStr, varietieId)

	return
}

func GetCustomerInfoByInfo(info string) (data []*models.Customer, err error) {
	sqlStr := `SELECT user_id, CONCAT(IFNULL(user_name,''),IFNULL(user_phone,'')) as info
				FROM user 
				where CONCAT(IFNULL(user_name,''),IFNULL(user_phone,''))
				like ?`
	err = db.Select(&data, sqlStr, "%"+info+"%")
	return
}

func GetCustomerInfoById(id int64) (info *models.UserInfo, err error) {
	sqlStr := `select user_name,user_phone,user_balance from user where user_id=?`
	info = new(models.UserInfo)
	err = db.Get(info, sqlStr,id)
	return
}
