package apps

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 所有管理员展示
func BalanceList(c *gin.Context) {
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("limit"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, n, err := models.BalanceList(page, pageSize, filters...)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err.Error(),
			"data":    "123",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "SUCCESS",
			"results":      list,
			"count":     n,
		})
	}
}

// 添加管理员
func AddBalance(c *gin.Context)  {
	b := new(models.Balance)
	b.Balance = c.Request.FormValue("balance")
	b.TransferIn = c.Request.FormValue("transfer_in")
	b.TransferOut = c.Request.FormValue("transfer_out")
	b.Interest = c.Request.FormValue("interest")
	b.Member = c.Request.FormValue("member")
	b.IsCompute, _ = strconv.Atoi(c.Request.FormValue("is_compute"))
	if id, err := b.AddBalance(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		b.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data": b,
		})
	}
}

// 删除管理员
func DeleteBalance(c *gin.Context) {
	bid, _ := strconv.Atoi(c.Request.FormValue("id"))
	if n, err := models.DeleteBalance(bid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}

// 更改管理员
func UpdateBalance(c *gin.Context)  {
	bid, _ := strconv.Atoi(c.Request.FormValue("id"))
	b := new(models.Balance)
	b.Id = bid
	b.Balance = c.Request.FormValue("balance")
	b.TransferIn = c.Request.FormValue("transfer_in")
	b.TransferOut = c.Request.FormValue("transfer_out")
	b.Interest = c.Request.FormValue("interest")
	b.Member = c.Request.FormValue("member")
	b.IsCompute, _ = strconv.Atoi(c.Request.FormValue("is_compute"))
	if n, err := b.UpdateBalance(bid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    n,
		})
	}
}
