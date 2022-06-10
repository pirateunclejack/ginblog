package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerAuthV1 := r.Group("/api/v1")
	routerAuthV1.Use(middleware.JwtToken())
	{
		// user mode router
		routerAuthV1.PUT("user/:id", v1.EditUser)
		routerAuthV1.DELETE("user/:id", v1.DeleteUser)
		routerAuthV1.GET("users", v1.GetUsers)

		// category mode router

		routerAuthV1.POST("category/add", v1.AddCate)
		routerAuthV1.PUT("category/:id", v1.EditCate)
		routerAuthV1.DELETE("category/:id", v1.DeleteCate)

		// article mode router

		routerAuthV1.POST("article/add", v1.AddArt)
		routerAuthV1.PUT("article/:id", v1.EditArt)
		routerAuthV1.DELETE("article/:id", v1.DeleteArt)

		// upload file
		routerAuthV1.POST("upload", v1.Upload)

	}

	routerPubV1 := r.Group("api/v1")
	{
		routerPubV1.POST("user/add", v1.AddUser)
		routerPubV1.GET("category", v1.GetCate)
		routerPubV1.GET("article", v1.GetArt)
		routerPubV1.GET("article/list/:cid", v1.GetCateArt)
		routerPubV1.GET("article/info/:id", v1.GetArtInfo)
		routerPubV1.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
