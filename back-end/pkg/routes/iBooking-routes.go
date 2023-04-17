package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/controllers"
	"net/http"
)

var RegisterBookingRoutes = func(router *gin.Engine) {
	// administrator management
	adminRouter := router.Group("/admin")
	{
		adminRouter.POST("/", controllers.CreateAdmin)
		adminRouter.POST("/login", controllers.AdminLogin)
	}

	// room management
	roomRouter := router.Group("/room")
	roomRouter.Use(controllers.UserAuthMiddleware.MiddlewareFunc())
	{
		roomRouter.GET("/", controllers.GetRoom)
		roomRouter.GET("/:roomID", controllers.GetRoomByID)
		auth := roomRouter.Group("/auth")
		auth.Use(controllers.AdminAuthMiddleware.MiddlewareFunc())
		{
			auth.GET("/", controllers.GetRoom)
			auth.GET("/:roomID", controllers.GetRoomByID)
			auth.POST("/", controllers.CreateRoom)
			auth.PUT("/", controllers.UpdateRoom)
			auth.DELETE("/:roomID", controllers.DeleteRoom)
		}
	}

	// seat management
	seatRouter := router.Group("/seat")
	seatRouter.Use(controllers.UserAuthMiddleware.MiddlewareFunc())
	{
		seatRouter.GET("/", controllers.GetSeat)
		seatRouter.GET("/:seatID", controllers.GetSeatByID)
		auth := seatRouter.Group("/auth")
		auth.Use(controllers.AdminAuthMiddleware.MiddlewareFunc())
		{
			auth.GET("/", controllers.GetSeat)
			auth.GET("/:seatID", controllers.GetSeatByID)
			auth.POST("/", controllers.CreateSeat)
			auth.PUT("/", controllers.UpdateSeat)
			auth.DELETE("/:seatID", controllers.DeleteSeat)
		}
	}

	// user management
	userRouter := router.Group("/user")
	{
		userRouter.POST("/", controllers.CreateUser)
		userRouter.POST("/login", controllers.UserLogin)
		auth := userRouter.Group("/auth")
		auth.Use(controllers.UserAuthMiddleware.MiddlewareFunc())
		{
			auth.POST("/logout", controllers.UserLogout)
			auth.POST("/refresh_token", controllers.UserRefreshToken)
			auth.PUT("/:userID", controllers.UpdateUser)
			auth.DELETE("/:userID", controllers.DeleteUser)
			auth.GET("/:userID", controllers.GetUserByID)
		}
	}

	// booking management
	//bookingRouter := router.Group("/booking")
	//{
	//	bookingRouter.POST("/", controllers.BookSeat)
	//	bookingRouter.GET("/", controllers.GetBook)
	//	bookingRouter.PUT("/:bookID", controllers.UpdateBook)    // update or attend
	//	bookingRouter.DELETE("/:bookID", controllers.DeleteBook) // cancel
	//}

	// notification management
	//notificationRouter := router.Group("/notification")
	//{
	//	notificationRouter.POST("/", controllers.CreateNotification)
	//}

	// default 404
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "template/404.html", nil)
	})
}
