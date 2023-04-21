package routes

import (
	"github.com/gin-gonic/gin"
	docs "github.com/littlema15/iBooking/docs" // main 文件中导入 docs 包
	"github.com/littlema15/iBooking/pkg/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var RegisterBookingRoutes = func(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			auth.GET("/", controllers.GetUserByUsername)
			auth.PUT("/password/:userID", controllers.UpdatePassword)
		}
	}

	// booking management
	bookingRouter := router.Group("/booking")
	{
		bookingRouter.POST("/", controllers.BookSeat)
		bookingRouter.GET("/getBookingByID/:bookingID", controllers.GetBookingByID)
		bookingRouter.GET("/:userID", controllers.GetBookingByUserID)
		bookingRouter.PUT("/:bookingID", controllers.UpdateBooking)    // update or attend
		bookingRouter.DELETE("/:bookingID", controllers.DeleteBooking) // cancel
	}

	// notification management
	notificationRouter := router.Group("/notification")
	{
		notificationRouter.POST("/", controllers.Notify)
	}

	// default 404
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "template/404.html", nil)
	})
}
