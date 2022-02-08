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
	public.POST("/ValidateUser/", Controller.ValidateUser)
	public.POST("/CheckTicket/", Controller.CheckTicket)
	public.POST("/User/Add", Controller.CreateUser)

	Claim := r.Group("/Claim")
	Claim.Use(Middleware.JwtAuthMiddlewareTicket())
	Claim.POST("/Ticket", Controller.ClaimTicket)

	admin := r.Group("/api/admin")
	admin.Use(Middleware.JwtAuthMiddleware())
	admin.GET("/Ticket", Controller.GetTicket)
	admin.GET("/Ticket/:id", Controller.GetTicketByID)
	admin.POST("/Ticket/Add", Controller.AddTicket)
	admin.PATCH("/Ticket/Delete/:id", Controller.DeleteTicket)
	admin.PATCH("/Ticket/Update/:id", Controller.UpdateTicket)
	admin.GET("/User", Controller.GetUser)
	admin.POST("/User/Add", Controller.CreateUser)
	admin.GET("/Ticket/UsedTicket", Controller.GetUsedTicket)
	admin.GET("/Ticket/UnusedTicket", Controller.GetUnusedTicket)
	admin.PATCH("/User/Delete/:id", Controller.DeleteUser)
	admin.PATCH("/User/Update/:id", Controller.UpdateUser)
	admin.PATCH("Ticket/RedeemTicket/:id", Controller.RedeemTicket)
	admin.GET("/Log", Controller.GetLog)
	admin.PATCH("User/ChangePassword/:id", Controller.ChangePassword)

	return r
}
