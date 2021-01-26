package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
)

func CommentRoute(routes *echo.Echo, commentApi api.CommentApi) {
	comment := routes.Group("/comment", IsAuthenticated)
	{
		comment.POST("/save", commentApi.SaveComment)
		comment.GET("/findAll/:tweet_id", commentApi.Comments)
	}
}