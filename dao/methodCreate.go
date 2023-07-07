package dao

import (
	"gin_demo/config"
	"gin_demo/models"
	"strconv"
	"sync"
)

// 根据Invoices切片插入对应发票数据信息
func PostInvs(invs []models.Invoice) {
	// 小于一万，直接插入
	if len(invs) < 10001 {
		config.Db.CreateInBatches(invs, 1000)
		// 此时记录数大于一万，采用协程方式插入
	} else {
		wg := sync.WaitGroup{}
		for i := 0; i < 6; i++ {
			wg.Add(1)
			go func(i int) {
				config.Db.CreateInBatches(invs[i*len(invs)/6:(i+1)*len(invs)/6], 1000)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	return
}

// 根据nums值生成对应个数的发票数据1
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

// 需求2：根据id生成手机号，id为1的手机号初始值为1001，id为2的手机号初始值为1002......
// 根据nums值生成对应个数的发票数据2
func CreateInvs1(nums int) (invs []models.Invoice) {
	invs = make([]models.Invoice, nums, nums)
	// 获得invs当前的最大id
	maxId := GetInvoicesMaxId()
	for i, j := 0, maxId+nums+1; i < nums; i, j = i+1, j+1 {
		invs[i].Number = "123"
		invs[i].InvoiceType = 1
		invs[i].Status = 1
		invs[i].ConsigneePhone = "100" + strconv.Itoa(j)
	}
	return
}

// 根据Invoices切片插入对应发票数据信息
func PutInvs(invs []models.Invoice) {
	config.Db.Debug().Create(invs)
	return
}
