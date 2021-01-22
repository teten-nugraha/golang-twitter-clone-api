package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/db"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"net/http"
)

func Init() *echo.Echo {
	db2 := db.InitDB()
	defer db2.Close()

	db2.AutoMigrate(domain.User{})
	db2.AutoMigrate(domain.Tweet{})
	db2.AutoMigrate(domain.TweetConversation{})

	routes := echo.New()
	routes.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return routes
}