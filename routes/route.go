package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teten-nugraha/golang-twitter-clone-api/db"
	"github.com/teten-nugraha/golang-twitter-clone-api/injection"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
	"net/http"
)

func Init() *echo.Echo {
	db2 := db.InitDB()
	//defer db2.Close()

	authApi := injection.InitAuthAPI(db2)
	userApi := injection.InitUserApi(db2)

	routes := echo.New()


	// Gzip Compression
	routes.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Allow CORS
	routes.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// set validator
	routes.Validator = &util.CustomValidator{Validator: validator.New()}

	routes.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	AuthRoute(routes, authApi)
	UserRoute(routes, userApi)

	return routes
}

// auth middleware
var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(util.SECRET),
})