package models

type Tabler interface {
	TableName() string
}

// ProductInformation 的表名重写为 `productInformation`
func (ProductInformation) TableName() string {
	return "productInformation"
}

// Consignee 的表名重写为 `consignee`
func (Consignee) TableName() string {
	return "consignee"
}
