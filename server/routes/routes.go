package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	RoomRoutes(e)
	UserRoutes(e)
	ChatRoutes(e)
}
