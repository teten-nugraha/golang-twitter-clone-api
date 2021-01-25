package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
)

func AuthRoute(routes *echo.Echo, authAPI api.AuthApi) {

	// auth route
	auth := routes.Group("/auth")
	{
		auth.GET("/generate-hash/:password", authAPI.GenerateHashPassword)
		auth.POST("/signup", authAPI.SignUp)
	}

}