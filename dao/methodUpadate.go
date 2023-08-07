package dao

import (
	"gin_demo/config"
	"gin_demo/models"
	"gorm.io/gorm/clause"
	"strconv"
	"sync"
	"time"
)

// 更新用户电话号码为特定格式，手机号格式：86+时间戳+编号
func UpadateUserPhone(nums int) {
	// 只需取出ID即可，这样能减少开销
	invsId := GetInvoicesId(nums)
	// 准备对应发票切片，该切片内的电话号码已经更新好至需要格式
	invs := make([]models.Invoice, nums)
	for i := 0; i < nums; i++ {
		invs[i].Id = int64(invsId[i])
		invs[i].ConsigneePhone = "86 + " + time.Now().Format("2006-01-02-15:04:05") + "+ 发票id：" + strconv.Itoa(invsId[i])
	}
	// 小于一千，单线程解决
	if nums < 1000 {
		config.Db.Debug().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"consignee_phone"}),
		}).Create(&invs)
		// 大于1000,开协程解决,
	} else {
		const routineNums = 10
		wg, numsFifth := sync.WaitGroup{}, nums/routineNums
		for i := 0; i < routineNums; i++ {
			wg.Add(1)
			go func(i int) {
				invsFifth := invs[numsFifth*i : numsFifth*(i+1)]
				config.Db.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "id"}},
					DoUpdates: clause.AssignmentColumns([]string{"consignee_phone"}),
				}).Create(&invsFifth)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	return
}
