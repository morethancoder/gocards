package main

import (
	"GoCards/internal/handler"
	"GoCards/views/pages"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

        e.Router.GET("/static/*", 
        apis.StaticDirectoryHandler(os.DirFS("./static"), false))

        e.Router.GET("/", func(c echo.Context) error {
            return handler.Render(c, http.StatusOK, pages.HomePage())
        })

        e.Router.GET("/dashboard", func(c echo.Context) error {
            return handler.Render(c, http.StatusOK, pages.DashboardPage())
        })

        e.Router.GET("/login", func(c echo.Context) error {
            return handler.Render(c, http.StatusOK, pages.LoginPage())
        })

        e.Router.GET("/signup", func(c echo.Context) error {
            return handler.Render(c, http.StatusOK, pages.SignupPage())
        })

        return nil
    })

    err := app.Start()
    if err != nil {
        log.Fatal(err)
    }
}
