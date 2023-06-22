package service

import (
	"fmt"
	"gin_demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetInvs(c *gin.Context) {
	// 获得size和page,此时都是string类型
	sizeS, pageS := c.Query("size"), c.Query("page")
	// 将size和page转为int类型
	size, err01 := strconv.Atoi(sizeS)
	if err01 != nil {
		fmt.Println("size输入有误：", err01)
	}
	page, err02 := strconv.Atoi(pageS)
	if err02 != nil {
		fmt.Println("page输入有误：", err02)
	}

	// 根据size和page查询数据库记录
	inv := dao.InvoiceQuerys(page, size)
	//invs := inv.InvoiceQuery(page, size)
	// 返回记录
	c.JSON(http.StatusOK, gin.H{
		"invs”": inv,
	})
	return
}
