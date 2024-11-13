package handler

import (
	"GoCards/views/layouts"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func Render(c echo.Context, status int, t templ.Component) error {
    //response header
    c.Response().Writer.WriteHeader(status)
    //response header content type
    c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

    isAlpine := c.Request().Header.Get("X-Requested-With") == "AlpineJS"
    route := c.Path()

    if !isAlpine {
        return layouts.BaseLayout(route).
        Render(c.Request().Context(), c.Response().Writer)
    }

    //render our templ component
    return t.Render(c.Request().Context(), c.Response().Writer)
}
