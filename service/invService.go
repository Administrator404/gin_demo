package service

import (
	"gin_demo/dao"
	"gin_demo/models"
)

func GetInvs(page int, size int) (invs []models.Invoice) {
	invs = dao.InvoiceQuerys(page, size)
	return invs
}
