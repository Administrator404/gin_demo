package dao

import (
	"gin_demo/config"
	"gin_demo/models"
)

// 根据ID，获取一个发票记录
func InvoiceQueryID(id int) (inv models.Invoice) {
	config.Db.Debug().Find(&inv, id)
	return inv
}

// 根据page和size获得对应页面的记录
func InvoiceQuerys(page int, size int) (invs []models.Invoice) {
	config.Db.Debug().Limit(size).Offset(page * size).Find(&invs)
	return
}

//// 根据发票号码获得发票的详情信息
//func InvoiceDetailsQueryNumber(number int) (invDetails []InvoiceDetails) {
//	dao.Db.Debug().Find(&invDetails, number)
//	fmt.Println(invDetails)
//	return invDetails
//}
