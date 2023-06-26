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

// 根据Invoices切片插入对应发票数据信息
func PutInvs(invs []models.Invoice) {
	config.Db.Debug().Create(invs)
	return
}

// 根据nums值生成对应个数的发票数据
func CreateInvs(nums int) (invs []models.Invoice) {
	invs = make([]models.Invoice, nums, nums)
	for i := 0; i < nums; i++ {
		invs[i].Number = "13"
		invs[i].InvoiceType = 1
		invs[i].Status = 1
		invs[i].ConsigneePhone = "13016179796"
	}
	return
}
