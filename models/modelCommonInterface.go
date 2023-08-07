package models

// Invs的常见方法
type Invs interface {
	CreateRecord() error
	DeleRecord() error
}
