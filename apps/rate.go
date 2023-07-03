package apps

import (
	"../models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 所有管理员展示
func RateList(c *gin.Context) {
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
	list, n, err := models.RateList(page, pageSize, filters...)
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
func AddRate(c *gin.Context)  {
	r := new(models.Rate)
	r.InterestRate = c.Request.FormValue("interest_rate")
	if id, err := r.AddRate(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
		})
	} else {
		r.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data": r,
		})
	}
}

// 删除管理员
func DeleteRate(c *gin.Context) {
	rid, _ := strconv.Atoi(c.Request.FormValue("id"))
	if n, err := models.DeleteRate(rid); err != nil {
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
func UpdateRate(c *gin.Context)  {
	rid, _ := strconv.Atoi(c.Request.FormValue("id"))
	r := new(models.Rate)
	r.Id = rid
	r.InterestRate = c.Request.FormValue("interest_rate")
	if n, err := r.UpdateRate(rid); err != nil {
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
