package service

import (
	"fmt"
	"gin_demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 根据page和size得到对应发票信息
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
}

// 根据发票号码得到发票信息、收货人信息、商品信息
func GetInvsALL(c *gin.Context) {
	// 获得发票编号
	invNumber := c.Param("invNumber")
	// 获得发票信息，收货人信息、商品信息
	inv := dao.InvoiceQuerysNumber(invNumber)
	product := dao.ProductQuerysByInvID(inv.Id)
	consignee := dao.ConsigneeQuerysByPhone(inv.ConsigneePhone)
	// 将信息返回给前端
	c.JSON(http.StatusOK, gin.H{
		"inv":       inv,
		"product":   product,
		"consignee": consignee,
	})
}
