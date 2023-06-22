package models

// 发票对应结构体
type Invoice struct {
	Id             int64  `json:"id"`
	Number         string `json:"number"`
	InvoiceType    uint8  `json:"type"`
	Status         uint8  `json:"status"`
	ConsigneePhone string `json:"consignee_phone"`
}

// 收货人
type Consignee struct {
	Name  string `json:"name"`
	Phone string `json:"phone",gorm:"primaryKey"`
}

// 商品信息
type ProductInformation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// 发票、商品表
type InvProduct struct {
	InvId     int64  `json:"inv_id"`
	ProductId string `json:"product_id"`
}
