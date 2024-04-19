package routes

import (
	"belajar-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func StudentRoute(e *echo.Echo) {
	e.POST("/student", controller.CreateStudent)
}