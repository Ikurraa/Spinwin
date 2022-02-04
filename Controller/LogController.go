package Controller

import (
	"awesomespinner/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetLog(c *gin.Context) {
	var log []Models.ViewLog
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&log)
	if len(log) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tidak ada log",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": log,
		})
	}
}
