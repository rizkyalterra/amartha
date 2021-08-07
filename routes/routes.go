package routes

import (
	"pembayaran/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	// e.GET("/users/:userId", DetailUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.POST("/users/register", controllers.RegisterUserController)
	return e
}
