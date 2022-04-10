package workout_scheduler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", schedule)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func hello_world(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
