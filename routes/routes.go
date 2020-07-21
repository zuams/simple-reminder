package routes

import (
	"github.com/zuams/simple-reminder/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	pfx := e.Group("/api")
	pfx.GET("/notes", controllers.GetNotes)
	pfx.POST("/notes", controllers.PostNote)
	pfx.PUT("/notes/:id", controllers.PutNote)
	pfx.DELETE("/notes/:id", controllers.DeleteNote)

	return e
}
