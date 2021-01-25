package api

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
	"net/http"
)

type TweetApi struct {
	TweetService service.TweetService
}

func ProviderTweetApi(u service.TweetService) TweetApi {
	return TweetApi{
		TweetService: u,
	}
}

func (t *TweetApi) SaveTweet(c echo.Context) error {

	form := new(dto.NewTweetDto)
	c.Bind(form)

	currentUserLogin := util.CurrentUserLogin(c)

	newTweetDto := dto.NewTweetDto{
		Tweet: form.Tweet,
		Maker: currentUserLogin.Username,
	}

	if err := c.Validate(form); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}



	tweetDto, err := t.TweetService.SaveTweet(newTweetDto)
	if err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return SuccessResponse(c, http.StatusOK, tweetDto)
}
