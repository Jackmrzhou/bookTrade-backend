package router

import (
	"bookTrade-backend/app/controllers"
	"bookTrade-backend/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) error {
	api := r.Group("/api")
	api.Use(middlewares.WrapResponse())
	imgAPI := api.Group("/image")
	{
		imgAPI.POST("/", controllers.UploadImage)
		imgAPI.GET("/", controllers.DownloadImage)
	}

	passportAPI := api.Group("/passport")
	{
		passportAPI.POST("/register", controllers.CreateNewAccount)
		passportAPI.POST("/verifyCode", controllers.SendVerifyCode)
		passportAPI.POST("/login", controllers.Login)
		passportAPI.GET("/logout", middlewares.AuthMiddleware(), controllers.Logout)
		passportAPI.GET("/ping", middlewares.AuthMiddleware(), controllers.Ping)
	}

	userAPI := api.Group("/user", middlewares.AuthMiddleware())
	{
		userAPI.POST("/", controllers.UpdateUser)
		userAPI.PUT("/", controllers.UpdateUser)
		userAPI.GET("/avatar", controllers.GetAvatar)
		userAPI.GET("/profile", controllers.GetUserProfile)
	}
	api.GET("/user/common", controllers.GetUserProfileCommon)
	//api.GET("/api/user", controllers.GetViewedUsers)

	bookAPI := api.Group("/book", middlewares.AuthMiddleware())
	{
		bookAPI.POST("/sell", controllers.NewSellBook)
		bookAPI.POST("/request", controllers.NewRequestBook)
	}
	api.GET("/book", controllers.GetBookDetail)
	api.GET("/book/viewed", controllers.GetViewedUsers)
	api.GET("/bookInOrder", controllers.GetBookDetailInOrder)

	catalogAPI := api.Group("/catalog")
	{
		catalogAPI.GET("/", controllers.GetCatalogs)
	}

	msgAPI := api.Group("/message", middlewares.AuthMiddleware())
	{
		msgAPI.GET("/contact", controllers.GetContact)
		msgAPI.POST("/contact", controllers.CreateContact)

		msgAPI.GET("/", controllers.GetAllMessage)
		msgAPI.POST("/", controllers.SendMessage)
		msgAPI.GET("/count", controllers.GetUnreadMsgCount)
	}

	listApi := api.Group("/list")
	{
		listApi.GET("/", controllers.ListAll)
	}

	orderAPI := api.Group("/order", middlewares.AuthMiddleware())
	{
		orderAPI.POST("/", controllers.CreateOrder)
		orderAPI.GET("/", controllers.GetOrders)
		orderAPI.PUT("/", controllers.UpdateOrder)
	}

	return nil
}
