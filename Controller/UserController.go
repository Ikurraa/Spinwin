package Controller

import (
	"awesomespinner/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	var user []Models.ViewUser
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&user)
	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tidak ada data admin",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func CreateUser(c *gin.Context) {
	var input Models.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	db := c.MustGet("db").(*gorm.DB)
	checkuser := db.Where("username = ?", input.Username).Limit(1).Find(&input)
	exist := checkuser.RowsAffected

	if exist == 0 {
		user := Models.CreateUser{Username: input.Username, Password: string(hashed_password), Role: "admin"}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    user,
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"Error": "Username Already Exist",
		})
		return
	}
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user Models.DeleteUser

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var input Models.DeleteTicket
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		passdata := Models.DeleteUser{
			Status: input.Status}
		db.Model(&user).Update(passdata)
		c.JSON(http.StatusOK, gin.H{
			"message": "User Berhasil dihapus",
		})
	}
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user Models.UpdateUser

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var input Models.InputUpdateUser
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var sameuser Models.UpdateUser
		if err := db.Where("username = ?", input.Username).Limit(1).Find(&sameuser).RowsAffected; err != 0 {
			c.JSON(http.StatusConflict, gin.H{
				"error": "username sudah terdaftar",
			})
		} else {
			passdata := Models.UpdateUser{
				Username: input.Username}
			db.Model(&user).Update(passdata)
			c.JSON(http.StatusOK, gin.H{
				"message": "user berhasil di update",
			})
		}
	}
}

func ValidateUser(c *gin.Context) {
	var input Models.InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	var user Models.ValidateUser
	if err := db.Where("username =?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			if user.Status == 0 {
				if user.Role == "Admins" {
					ValidToken, err := GenerateJWTUser(user.ID, user.Role)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"error": err.Error(),
						})
						return
					} else {
						timenow := time.Now().Format(time.RFC3339)
						login := Models.ValidateUser{
							Last_login: timenow}
						db.Model(&user).Update(login)
						c.JSON(http.StatusOK, gin.H{
							"token":     ValidToken,
							"message":   1,
							"user_role": user.Role,
						})
					}
				} else {
					ValidToken, err := GenerateJWTUser(user.ID, user.Role)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"error": err.Error(),
						})
						return
					} else {
						timenow := time.Now().Format(time.RFC3339)
						login := Models.ValidateUser{
							Last_login: timenow}
						db.Model(&user).Update(login)
						c.JSON(http.StatusOK, gin.H{
							"token":     ValidToken,
							"message":   0,
							"user_role": user,
						})
					}
				}

			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "User tidak terdaftar",
				})
				return
			}
		}
	}
}

func GetCurrentUser(c *gin.Context) (Models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	user_id, err := ExtractTokenIDUser(c)
	if err != nil {
		return Models.User{}, err
	}
	var user Models.User
	if err := db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return user, fmt.Errorf("User Tidak ditemukan %s ", user_id)
	} else {
		return user, nil
	}
}

func GetCurrentRoleUser(c *gin.Context) {
	role, err := ExtractTokenRoleUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Terjadi kesalahan",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"role": role,
		})
	}
}

func ChangePassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user Models.ChangePassword
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		var input Models.ChangePassword
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			hashed_password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
			passdata := Models.ChangePassword{
				Password: string(hashed_password)}
			db.Model(&user).Update(passdata)
			c.JSON(http.StatusOK, gin.H{
				"message": "Sukses mengubah password",
			})
		}
	}
}
