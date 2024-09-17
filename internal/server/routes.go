package server

import (
	"net/http"

	"dictionary-htmx-demo/cmd/web"
	"dictionary-htmx-demo/cmd/web/list"
	"dictionary-htmx-demo/cmd/web/test"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/listPage", echo.WrapHandler(templ.Handler(list.ListPage())))
	e.GET("/list", echo.WrapHandler(templ.Handler(list.List())))
	e.GET("/test", echo.WrapHandler(http.HandlerFunc(test.TestTabHandler)))

	// e.GET("/test", echo.WrapHandler(templ.Handler(test.Test())))

	// e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	// e.GET("/", s.HelloWorldHandler)

	// e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

// func (s *Server) healthHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, s.db.Health())
// }
