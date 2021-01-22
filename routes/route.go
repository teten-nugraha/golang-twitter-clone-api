package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/teten-nugraha/golang-twitter-clone-api/db"
	"github.com/teten-nugraha/golang-twitter-clone-api/injection"
	"net/http"
)

func Init() *echo.Echo {
	db2 := db.InitDB()
	//defer db2.Close()

	userApi := injection.InitUserApi(db2)

	routes := echo.New()
	routes.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	UserRoute(routes, userApi)

	return routes
}