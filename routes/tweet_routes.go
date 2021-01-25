package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/api"
)

func TweetRoute(routes *echo.Echo, tweetApi api.TweetApi) {
	tweet := routes.Group("/tweet", IsAuthenticated)
	{
		tweet.POST("/save", tweetApi.SaveTweet)
	}

}
