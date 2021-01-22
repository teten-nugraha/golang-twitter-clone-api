package main

import (
	"github.com/teten-nugraha/golang-twitter-clone-api/routes"
	"github.com/teten-nugraha/golang-twitter-clone-api/util"
)

func main() {
	route := routes.Init()

	route.Logger.Fatal(route.Start(util.PORT))
}
