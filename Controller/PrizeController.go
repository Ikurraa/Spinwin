package Controller

import (
	"awesomespinner/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetPrize(c *gin.Context) {
	var prize []Models.Prize
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&prize)
	if len(prize) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data tidak ditemukan",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    prize,
			"message": "success",
		})
	}
}

func AddPrize(c *gin.Context) {
	var input Models.InputAddPrize
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		db := c.MustGet("db").(*gorm.DB)
		var prize Models.Prize
		if err := db.Where("name = ?", input.Name).Limit(1).First(&prize).RowsAffected; err != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Hadiah sudah terdaftar",
			})
			return
		} else {
			data := Models.AddPrize{
				Name:       input.Name,
				Value:      input.Value,
				Percentage: input.Percentage,
			}
			db.Create(&data)
			c.JSON(http.StatusOK, gin.H{
				"message": "hadiah berhasil ditambahkan",
			})
		}
	}
}

func UpdatePrize(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var prize Models.Prize
	if err := db.Where("id = ?", c.Param("id")).First(&prize).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var input Models.InputUpdatePrize
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			data := Models.UpdatePrize{
				Name:       input.Name,
				Value:      input.Value,
				Percentage: input.Percentage,
			}
			db.Model(&prize).Update(data)
			c.JSON(http.StatusOK, gin.H{
				"message": "Hadiah berhasil di update",
			})
		}
	}
}

func DeletePrize(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var prize Models.Prize
	if err := db.Where("id = ?", c.Param("id")).First(&prize).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		db.Delete(&prize)
		c.JSON(http.StatusOK, gin.H{
			"message": "hadiah berhasil dihapus",
		})
	}
}

func GetPrizeList(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var prize []Models.Prize
	db.Find(&prize)
	//c.JSON(http.StatusOK, gin.H{
	//	"data": prize,
	//})
	var input Models.GetPrizeNum
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var a int
	var list []Models.Prize
	for i, prize := range prize {
		if i == 0 {
			a = prize.Percentage
			list = append(list, prize)
		} else {
			prize.Percentage = prize.Percentage + a
			list = append(list, prize)
			a = prize.Percentage
		}
	}
	var g int
	if input.Number > 100 {
		input.Number = 1
	}
	for i, list := range list {
		if i == 0 {
			if input.Number >= 0 && input.Number <= list.Percentage {
				c.JSON(http.StatusOK, gin.H{
					"data": list,
				})
				g = list.Percentage
			}
		} else {
			if input.Number > g && input.Number <= list.Percentage {
				c.JSON(http.StatusOK, gin.H{
					"data": list,
				})
				g = list.Percentage
			}
		}
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"data": list,
	//})
}
