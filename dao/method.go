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

// 根据发票号码查询发票信息
func InvoiceQuerysNumber(invNumber string) (inv models.Invoice) {
	config.Db.Debug().Where("number = ?", invNumber).Find(&inv)
	return
}

// 根据发票ID查询商品信息
func ProductQuerysByInvID(invId int64) (product models.ProductInformation) {
	config.Db.Debug().Where("id = ?", invId).First(&product)
	return
}

// 根据收货人号码查询收货人信息
func ConsigneeQuerysByPhone(consigneePhone string) (consignee models.Consignee) {
	config.Db.Debug().Where("phone = ?", consigneePhone).First(&consignee)
	return
}
