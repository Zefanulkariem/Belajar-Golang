package main

import (
	"belajar-rest-api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Amerika Yaa!")
	})

	routes.StudentRoute(e)

	e.Logger.Fatal(e.Start(":1323"))
}

