package server

import (
	"net/http"

	"dictionary-htmx-demo/cmd/web"
	"dictionary-htmx-demo/cmd/web/list"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/get-list-page", echo.WrapHandler(http.HandlerFunc(list.ListPageHandler)))
	e.GET("/switch-to-list", echo.WrapHandler(http.HandlerFunc(list.ListHandler)))
	e.GET("/get-list-content", echo.WrapHandler(http.HandlerFunc(list.GetListContentHandler)))
	e.GET("/get-add-word-form", echo.WrapHandler(http.HandlerFunc(list.GetAddWordFormHandler)))
	e.POST("/add-word", echo.WrapHandler(http.HandlerFunc(list.AddWordHandler)))
	e.GET("/get-update-word-form", echo.WrapHandler(http.HandlerFunc(list.GetUpdateWordFormHandler)))
	e.PUT("/update-word", echo.WrapHandler(http.HandlerFunc(list.UpdateWordHandler)))
	e.DELETE("/delete-word", echo.WrapHandler(http.HandlerFunc(list.DeleteWord)))
	e.DELETE("/delete-component", s.EmptyHandler)

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

func (s *Server) EmptyHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}
