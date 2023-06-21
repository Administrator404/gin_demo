package models

// 发票对应结构体
type Invoice struct {
	Id          int64  `json:"id"`
	Number      string `json:"number"`
	InvoiceType uint8  `json:"type"`
	Status      uint8  `json:"status"`
}

// 收货人
type Consignee struct {
	Name  string `json:"name"`
	phone string `json:"phone",gorm:"primaryKey"`
}

// 商品信息
type ProductInformation struct {
	id   string `json:"id"`
	name string `json:"name"`
}
