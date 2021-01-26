package api

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/dto"
	"github.com/teten-nugraha/golang-twitter-clone-api/service"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
	"net/http"
	"strconv"
)

type CommentApi struct {
	CommentService service.CommentService
}

func ProviderCommentApi(u service.CommentService) CommentApi {
	return CommentApi{
		CommentService: u,
	}
}

func (c *CommentApi) SaveComment(e echo.Context) error {
	payload := new(dto.CommentDto)
	if err := e.Bind(payload); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	if err := e.Validate(payload); err != nil {
		return ErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	currentUserLogin := util.CurrentUserLogin(e)

	commentDto := dto.CommentDto{
		TweetId: payload.TweetId,
		Maker: currentUserLogin.Username,
		Reply: payload.Reply,
	}

	commentDto, err := c.CommentService.SaveComment(commentDto)
	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, commentDto)
}

func (c *CommentApi) Comments(e echo.Context) error {
	tweet_id, err := strconv.ParseUint(string(e.Param("tweet_id")), 10, 64)
	if err != nil {
		return ErrorResponse(e, http.StatusBadRequest, err.Error())
	}

	commentsDtos := c.CommentService.Comments(tweet_id)

	return SuccessResponse(e, http.StatusOK, commentsDtos)
}