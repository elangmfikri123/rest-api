package routes

import (
	"github.com/elangmfikri123/rest-api/app/user"
	"github.com/elangmfikri123/rest-api/helper"
	"github.com/labstack/echo/v4"
)

func DefineApiRoutes(e *echo.Echo) {
	handler := []helper.Handler{
		user.UserRoutes{},
	}

	var routes []helper.Route

	for _, handler := range handler {
		routes = append(routes, handler.Route()...)
	}

	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				e.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				e.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				e.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				e.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				e.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		}
	}
}
