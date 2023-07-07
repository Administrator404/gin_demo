package models

import "gin_demo/config"

// 创建对应记录
func (this *Invoice) CreateRecord() error {
	err := config.Db.Create(&this).Error
	return err
}

// 删除对应记录
func (this *Invoice) DeleRecord() error {
	err := config.Db.Delete(&this).Error
	return err
}
