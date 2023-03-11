package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sunny092020/go-master-data/controllers/middleware"
	"github.com/sunny092020/go-master-data/controllers/v1"
	"github.com/sunny092020/go-master-data/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin)
			auth.POST("/signup", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1.GETUser)
			user.PUT("", utils.AuthOnly, v1.PUTUser)
		}

		ping := v1route.Group("/ping")
		{
			ping.GET("", v1.GETPing)
		}
	}
	return
}
