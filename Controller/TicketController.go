package Controller

import (
	"awesomespinner/Models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func GetTicket(c *gin.Context) {
	var ticket []Models.ViewTicket
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&ticket)
	if len(ticket) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tidak ada ticket",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": ticket,
		})
	}
}

func GetUsedTicket(c *gin.Context) {
	var ticket []Models.ViewUsedTicket
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&ticket)
	if len(ticket) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tidak ada ticket",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": ticket,
		})
	}
}

func GetUnusedTicket(c *gin.Context) {
	var ticket []Models.ViewUnusedTicket
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&ticket)
	if len(ticket) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tidak ada ticket",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": ticket,
		})
	}
}

func GetTicketByID(c *gin.Context) {
	var ticket Models.ViewTicket
	db := c.MustGet("db").(*gorm.DB)
	err := db.First(&ticket, "id = ?", c.Param("id")).Find(&ticket).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tidak ada ticket",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": ticket,
		})
	}
}

func AddTicket(c *gin.Context) {
	var input Models.InputTicket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  user,
		})
		return
	}
	formticket := Models.AddTicket{
		Ticket_code:   input.Ticket_code,
		Player_name:   input.Player_name,
		Ticket_status: 0,
		Redeem_status: 0,
		Created_by:    user.Username}
	db := c.MustGet("db").(*gorm.DB)
	check := db.Where("ticket_code = ?", input.Ticket_code).Limit(1).Find(&formticket)
	if err := check.RowsAffected; err != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nomor ticket sudah ada",
		})
		return
	} else {
		log := Models.Log{
			User_id:     user.ID,
			Last_update: "Create ticket " + formticket.Ticket_code,
		}
		db.Create(&formticket)
		db.Create(&log)
		c.JSON(http.StatusOK, gin.H{
			"message": "input ticket berhasil",
			"data":    formticket,
		})
	}
}

func DeleteTicket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ticket Models.DeleteTicket

	if err := db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data tidak ditemukan",
		})
		return
	} else {
		var input Models.InputDeleteTicket
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := GetCurrentUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		timenow := time.Now().Format(time.RFC3339)
		passdata := Models.DeleteTicket{
			Status:     input.Status,
			Update_by:  user.Username,
			Updated_at: timenow}
		log := Models.Log{
			User_id:     user.ID,
			Last_update: "Delete Ticket " + ticket.Ticket_code}
		db.Model(&ticket).Update(passdata)
		db.Create(&log)
		c.JSON(http.StatusOK, gin.H{
			"message": "Ticket Berhasil dihapus",
			"ticket":  ticket,
		})
	}
}

func UpdateTicket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ticket Models.UpdateTicket

	if err := db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data tidak ditemukan",
		})
		return
	} else {
		var input Models.InputUpdateTicket
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, err := GetCurrentUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			if ticket.Status == 0 {
				if ticket.Ticket_status == 0 {
					user := Models.User{
						ID:       user.ID,
						Username: user.Username}
					timenow := time.Now().Format(time.RFC3339)
					passdata := Models.UpdateTicket{
						Ticket_code: input.Ticket_code,
						Player_name: input.Player_name,
						Update_by:   user.Username,
						Updated_at:  timenow}
					log := Models.Log{
						User_id:     user.ID,
						Last_update: "Update Ticket " + passdata.Ticket_code}
					db.Model(&ticket).Update(passdata)
					db.Create(&log)
					c.JSON(http.StatusOK, gin.H{
						"message": "Ticket Berhasil diupdate",
					})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "Ticket tidak terdaftar",
					})
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Ticket tidak terdaftar",
				})
				return
			}
		}
	}

}

func CheckTicket(c *gin.Context) {
	var input Models.InputCheckTicket
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ticket Models.CheckTicket
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("ticket_code = ? and player_name = ?", input.Ticket_Code, input.Player_name).First(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ticket tidak ada",
		})
		return
	} else {
		if ticket.Ticket_status == 0 && ticket.Status == 0 {
			ValidToken, err := GenerateJWTTicket(ticket.Id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"token":   ValidToken,
					"message": "ticket dapat digunakan",
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ticket tidak dapat digunakan",
			})
		}
	}
}

func ClaimTicket(c *gin.Context) {
	var InputReward Models.InputReward
	if err := c.ShouldBindJSON(&InputReward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	reward := InputReward.Reward
	value := InputReward.Value
	ticket_id, err := ExtractTokenIDTicket(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var ticket Models.ClaimTicket
	ticket.Id = ticket_id
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", ticket.Id).First(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		if ticket.Ticket_status == 0 && ticket.Status == 0 {
			timenow := time.Now().Format(time.RFC3339)
			passticket := Models.ClaimTicket{
				Ticket_status: 1,
				Redeem_at:     timenow}
			db.Model(&ticket).Update(passticket)
			passreward := Models.Reward{
				ID:     ticket.Ticket_code,
				Reward: reward,
				Value:  value,
			}
			db.Create(&passreward)
			c.JSON(http.StatusOK, gin.H{
				"ticket":    ticket,
				"reward":    passreward,
				"isireward": InputReward.Reward,
				"Message":   "Ticket berhasil di claim",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Ticket tidak dapat digunakan",
			})
			return
		}
	}
}

func RedeemTicket(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var ticket Models.RedeemTicket
	if err := db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var input Models.InputTicket
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if input.Ticket_code == ticket.Ticket_code {
			if ticket.Ticket_status == 1 && ticket.Status == 0 {
				timenow := time.Now().Format(time.RFC3339)
				user, err := GetCurrentUser(c)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}
				passdata := Models.RedeemTicket{
					Id:            ticket.Id,
					Ticket_code:   input.Ticket_code,
					Update_by:     user.Username,
					Updated_at:    timenow,
					Redeem_status: 1}
				log := Models.Log{
					User_id:     user.ID,
					Last_update: "Redeem Ticket " + ticket.Ticket_code}
				db.Model(&ticket).Update(passdata)
				db.Create(&log)
				c.JSON(http.StatusOK, gin.H{
					"message": "Ticket sudah di redeem",
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Ticket tidak dapat di redeem",
				})
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Terjadi kesalahan",
			})
			return
		}
	}
}
