package Routes

import (
	"awesomespinner/Controller"
	"awesomespinner/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(Middleware.CORS())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	public := r.Group("/api")
	Claim := r.Group("/Claim")
	admin := r.Group("/api/admin")
	admin.Use(Middleware.JwtAuthMiddleware())
	Claim.Use(Middleware.JwtAuthMiddlewareTicket())

	//Ticket
	public.POST("/CheckTicket/", Controller.CheckTicket)
	Claim.POST("/Ticket", Controller.ClaimTicket)
	admin.GET("/Ticket", Controller.GetTicket)
	admin.GET("/Ticket/:id", Controller.GetTicketByID)
	admin.POST("/Ticket/Add", Controller.AddTicket)
	admin.PATCH("/Ticket/Delete/:id", Controller.DeleteTicket)
	admin.PATCH("/Ticket/Update/:id", Controller.UpdateTicket)
	admin.GET("/Ticket/UsedTicket", Controller.GetUsedTicket)
	admin.GET("/Ticket/UnusedTicket", Controller.GetUnusedTicket)
	admin.PATCH("Ticket/RedeemTicket/:id", Controller.RedeemTicket)

	//User
	public.POST("/ValidateUser/", Controller.ValidateUser)
	public.POST("/User/Add", Controller.CreateUser)
	admin.GET("/User", Controller.GetUser)
	admin.POST("/User/Add", Controller.CreateUser)
	admin.PATCH("/User/Delete/:id", Controller.DeleteUser)
	admin.PATCH("/User/Update/:id", Controller.UpdateUser)
	admin.PATCH("User/ChangePassword/:id", Controller.ChangePassword)
	admin.GET("User/GetRole", Controller.GetCurrentRoleUser)

	//Log
	admin.GET("/Log", Controller.GetLog)

	//Prize
	public.GET("/Prize/List", Controller.GetPrizeList)
	public.POST("/Prize/GetValueList", Controller.GetPrizeReward)
	admin.GET("/Prize/GetPrize", Controller.GetPrize)
	admin.POST("/Prize/AddPrize", Controller.AddPrize)
	admin.PATCH("/Prize/UpdatePrize/:id", Controller.UpdatePrize)
	admin.DELETE("/Prize/DeletePrize/:id", Controller.DeletePrize)

	return r
}
